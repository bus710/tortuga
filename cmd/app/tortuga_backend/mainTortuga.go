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
	ticker := time.NewTicker(100 * time.Millisecond).C

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

// calculate translates the button coordinates (x, y)
// to the Kobuki command (speed, angle)
func (t *Tortuga) calculate() {
	// http://yujinrobot.github.io/kobuki/enAppendixProtocolSpecification.html

	if t.request == t.last {
		// If the current request (request) is same as the last request sent to the robot,
		// don't need to stop or change the angle but increase the speed to the max as LUT progressivly
		t.request = t.last // not really need
		// TODO: check the speed and angle with the LUT and increase the speed
	} else {
		// If the current request (request) is not same as the last request sent to the robot,
		// need to ignore the request for now but decrease the speed to 0 progressively
		if t.speed == 0 && t.angle == 0 {
			t.request = t.last
		} else {
		}
	}
	// speedAngle := t.lut[t.request.y][t.request.x]
	// t.speed = int16(speedAngle[0])
	// t.angle = int16(speedAngle[1])

}
