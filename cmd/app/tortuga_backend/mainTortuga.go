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
	current         BasicControl
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
	t.current = BasicControl{"none", "none"}

	t.speed = 0
	t.angle = 0

	t.timeOverCounter = 0

	// https://www.tutorialspoint.com/go/go_multi_dimensional_arrays.htm
	t.lut = [3][3][2]int{
		{{100, 100}, {100, 0}, {100, -100}},
		{{100, 1}, {0, 0}, {-100, 1}},
		{{-100, 100}, {-100, 0}, {-100, -100}},
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
	// speedAngle := t.lut[t.request.y][t.request.x]
	// t.speed = int16(speedAngle[0])
	// t.angle = int16(speedAngle[1])

	//===================================================
	// // if there is no input, stop the robot but smoothly
	// if t.request.x == 2 && t.request.y == 2 {
	// 	// if the previous action was the pure rotation
	// 	if t.current.DraggedX == 1 || t.current.DraggedX == -1 {
	// 		t.speed = 0
	// 		t.angle = 0
	// 		return
	// 	}
	// 	if t.current.DraggedY > 10 {
	// 		t.current.DraggedY -= 10
	// 	} else if t.current.DraggedY < -10 {
	// 		t.current.DraggedY += 10
	// 	} else {
	// 		t.current.DraggedX = 0
	// 		t.current.DraggedY = 0
	// 	}
	// 	t.speed = t.current.DraggedY
	// 	t.angle = 0
	// 	return
	// }
	//===================================================
	// // forward/backward calculation
	// if t.request.DraggedY != t.current.DraggedY {
	// 	// Calculate a new forward/backward movement value
	// 	// based on the	difference between the request and current
	// 	if t.request.DraggedY > t.current.DraggedY {
	// 		// if the request is further than the current
	// 		diff := t.request.DraggedY - t.current.DraggedY
	// 		if diff > 10 {
	// 			t.current.DraggedY += 10
	// 		} else {
	// 			t.current.DraggedY = t.request.DraggedY
	// 		}
	// 	} else {
	// 		if t.request.DraggedY > 0 {
	// 			// if the request is closer than the current (both are positive)
	// 			diff := t.current.DraggedY - t.request.DraggedY
	// 			if diff > 10 {
	// 				t.current.DraggedY -= 10
	// 			} else {
	// 				t.current.DraggedY = t.request.DraggedY
	// 			}
	// 		} else {
	// 			diff := math.Abs(float64(t.request.DraggedY)) - math.Abs(float64(t.current.DraggedY))
	// 			if diff > 10 {
	// 				t.current.DraggedY -= 10
	// 			} else {
	// 				t.current.DraggedY = t.request.DraggedY
	// 			}
	// 		}
	// 	}
	// }
	// t.speed = t.current.DraggedY
	// t.angle = t.current.DraggedX
	// return
	//===================================================
	// pure rotation
	// if (t.request.DraggedY < 10 && t.request.DraggedY > -10) &&
	// 	(t.request.DraggedX > 30 || t.request.DraggedX < -30) {
	// 	// limiter
	// 	if t.request.DraggedX > 50 {
	// 		t.request.DraggedX = 50
	// 	} else if t.request.DraggedX < -50 {
	// 		t.request.DraggedX = -50
	// 	}
	// 	t.current.DraggedX = 1
	// 	t.current.DraggedY = t.request.DraggedX * -1
	// 	t.speed = t.current.DraggedY
	// 	t.angle = t.current.DraggedX
	// 	return
	// }
	//===================================================
	// rotationRequest := int16(0)
	// rotationCurrent := int16(0)
	// if t.request.DraggedX > 0 {
	// 	rotationRequest = t.request.DraggedY * (t.request.DraggedX + 230/2) / t.request.DraggedX
	// 	rotationCurrent = t.current.DraggedY * (t.current.DraggedX + 230/2) / t.request.DraggedX
	// 	diff := math.Abs(float64(rotationRequest)) - math.Abs(float64(rotationCurrent))
	// 	if diff > 10 {
	// 		if rotationRequest > rotationCurrent {
	// 			rotationCurrent -= 10
	// 		} else {
	// 			rotationCurrent += 10
	// 		}
	// 	}
	// } else if t.request.DraggedX < 0 {
	// 	rotationRequest = t.request.DraggedY * (t.request.DraggedX - 230/2) / t.request.DraggedX
	// 	rotationCurrent = t.current.DraggedY * (t.current.DraggedX - 230/2) / t.current.DraggedX
	// 	diff := math.Abs(float64(rotationRequest)) - math.Abs(float64(rotationCurrent))
	// 	if diff > 10 {
	// 		if rotationRequest > rotationCurrent {
	// 			rotationCurrent += 10
	// 		} else {
	// 			rotationCurrent -= 10
	// 		}
	// 	}
	// } else {
	// 	t.current.DraggedX = 0
	// 	rotationCurrent = 0
	// }
	// t.speed = t.current.DraggedY
	// t.angle = rotationCurrent
	// return
	//===================================================
}
