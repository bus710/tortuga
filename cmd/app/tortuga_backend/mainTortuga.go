package main

import (
	"log"
	"time"

	"github.com/bus710/tortuga"
	"github.com/bus710/tortuga/cmd/model"
)

// BasicControl ...
type BasicControl struct {
	forwardBackward string
	leftRight       string
}

// Tortuga ...
type Tortuga struct {
	app  *App
	conn tortuga.Connection

	chanStop    chan bool
	chanRequest chan BasicControl

	timeOverCounter int
	lut             [3][3][2]int
	request         BasicControl
	last            BasicControl
	speed           int16
	angle           int16

	battery byte
}

func (t *Tortuga) init(app *App) {

	t.app = app
	t.conn = tortuga.Connection{}
	t.chanStop = make(chan bool, 1)
	t.chanRequest = make(chan BasicControl, 1)

	t.request = BasicControl{"none", "none"}
	t.last = BasicControl{"none", "none"}

	t.speed = 0
	t.angle = 0

	t.timeOverCounter = 0

	// https://www.tutorialspoint.com/go/go_multi_dimensional_arrays.htm
	t.lut = [3][3][2]int{
		{{100, 50}, {100, 0}, {100, -50}},
		{{100, 1}, {0, 0}, {-100, 1}},
		{{-100, 50}, {-100, 0}, {-100, -50}},
	}

	err := t.conn.Init(&app.waitInstance, t.handler, "ttyUSB0")
	if err != nil {
		// log.Fatal(err)
	}
}

func (t *Tortuga) run() {

	go t.conn.Run()
	ticker := time.NewTicker(500 * time.Millisecond).C

run:
	for {
		select {
		case <-ticker:
			// if there is no request for 3 seconds,
			// the robot should stop
			t.timeOverCounter++
			if t.timeOverCounter > 30 {
				t.timeOverCounter = 0
				t.request = BasicControl{"none", "none"}
			}
			t.calculate()
			log.Println(t.request, t.speed, t.angle)
			// t.conn.Send(command.BaseControlCommand(t.speed, t.angle))

		case request := <-t.chanRequest:
			t.timeOverCounter = 0
			t.request = request
			log.Println(t.request.forwardBackward, "/", t.request.leftRight)

		case <-t.chanStop:
			t.conn.Stop()
			time.Sleep(time.Second)
			break run
		}
	}
	t.app.waitInstance.Done()
}

func (t *Tortuga) handler(fdb model.Feedback) {
	t.battery = fdb.BasicSensorData.Battery
}

// kind of motion planning
func (t *Tortuga) calculate() {
	// http://yujinrobot.github.io/kobuki/enAppendixProtocolSpecification.html

	if t.request == t.last {
		// If the current request (request) is same as the last request sent to the robot,
		// don't need to stop or change the angle but increase the speed to the max as LUT progressivly
		speedAngle := [2]int{0, 0}
		switch {
		case t.request.forwardBackward == "forward" && t.request.leftRight == "left":
			speedAngle = t.lut[0][0]
			if t.speed < int16(speedAngle[0]) {
				t.speed += 10
			}
			if t.angle < int16(speedAngle[1]) {
				t.angle += 10
			}
		case t.request.forwardBackward == "forward" && t.request.leftRight == "none":
			speedAngle = t.lut[0][1]
			if t.speed < int16(speedAngle[0]) {
				t.speed += 10
			}
			t.angle = 0
		case t.request.forwardBackward == "forward" && t.request.leftRight == "right":
			speedAngle = t.lut[0][2]
			if t.speed < int16(speedAngle[0]) {
				t.speed += 10
			}
			if t.angle > int16(speedAngle[1]) {
				t.angle -= 10
			}
		case t.request.forwardBackward == "none" && t.request.leftRight == "left":
			speedAngle = t.lut[1][0]
			if t.speed < int16(speedAngle[0]) {
				t.speed += 10
			}
			if t.angle < int16(speedAngle[1]) {
				t.angle = 1
			}
		case t.request.forwardBackward == "none" && t.request.leftRight == "none":
			t.speed = 0
			t.angle = 0
		case t.request.forwardBackward == "none" && t.request.leftRight == "right":
			speedAngle = t.lut[1][2]
			if t.speed > int16(speedAngle[0]) {
				t.speed -= 10
			}
			if t.angle < int16(speedAngle[1]) {
				t.angle = 1
			}
		case t.request.forwardBackward == "backward" && t.request.leftRight == "left":
			speedAngle = t.lut[2][0]
			if t.speed > int16(speedAngle[0]) {
				t.speed -= 10
			}
			if t.angle < int16(speedAngle[1]) {
				t.angle += 10
			}
		case t.request.forwardBackward == "backward" && t.request.leftRight == "none":
			speedAngle = t.lut[2][1]
			if t.speed > int16(speedAngle[0]) {
				t.speed -= 10
			}
			t.angle = 0
		case t.request.forwardBackward == "backward" && t.request.leftRight == "right":
			speedAngle = t.lut[2][2]
			if t.speed > int16(speedAngle[0]) {
				t.speed -= 10
			}
			if t.angle > int16(speedAngle[1]) {
				t.angle -= 10
			}
		default:
			t.speed = 0
			t.angle = 0
		}
	} else {
		// If the current request (request) is not same as the last request sent to the robot,
		// need to ignore the request for now but decrease the speed to 0 progressively
		if t.speed == 0 && t.angle == 0 {
			t.last = t.request
		} else {
			// If the last request stil affects to the robot (speed and angle are not 0),
			// decrease the params as much as 10 until 0
			if t.speed > 10 {
				t.speed -= 10
			} else if t.speed < -10 {
				t.speed += 10
			} else if t.speed == 10 {
				t.speed = 0
				t.angle = 0
			} else if t.speed == -10 {
				t.speed = 0
				t.angle = 0
			}
		}
	}
}
