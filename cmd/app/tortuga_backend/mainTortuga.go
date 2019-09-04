package main

import (
	"log"
	"math"
	"time"

	"github.com/bus710/tortuga"
	"github.com/bus710/tortuga/cmd/command"
	"github.com/bus710/tortuga/cmd/model"
)

// BasicControl ...
type BasicControl struct {
	OriginalX int16
	OriginalY int16
	DraggedX  int16
	DraggedY  int16
}

// Tortuga ...
type Tortuga struct {
	app         *App
	conn        tortuga.Connection
	chanStop    chan bool
	chanRequest chan BasicControl
	request     BasicControl
	current     BasicControl
	speed       int16
	angle       int16
	battery     byte
}

func (t *Tortuga) init(app *App) {

	t.app = app
	t.conn = tortuga.Connection{}
	t.chanStop = make(chan bool, 1)
	t.chanRequest = make(chan BasicControl, 1)

	t.speed = 0
	t.angle = 0

	err := t.conn.Init(&app.waitInstance, t.handler, "ttyUSB0")
	if err != nil {
		log.Fatal(err)
	}
}

func (t *Tortuga) run() {

	go t.conn.Run()

	ticker := time.NewTicker(100 * time.Millisecond).C

run:
	for {
		select {
		case <-ticker:
			t.calculate()
			t.conn.Send(command.BaseControlCommand(t.speed, t.angle))

		case request := <-t.chanRequest:
			t.request = request

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

func (t *Tortuga) calculate() {
	// http://yujinrobot.github.io/kobuki/enAppendixProtocolSpecification.html

	// if there is no input, stop the robot but smoothly
	if t.request.OriginalX == 0 && t.request.OriginalY == 0 &&
		t.request.DraggedX == 0 && t.request.DraggedY == 0 {

		// if the previous action was the pure rotation
		if t.current.DraggedX == 1 || t.current.DraggedX == -1 {
			t.speed = 0
			t.angle = 0
			return
		}

		if t.current.DraggedY > 10 {
			t.current.DraggedY -= 10
		} else if t.current.DraggedY < -10 {
			t.current.DraggedY += 10
		} else {
			t.current.DraggedX = 0
			t.current.DraggedY = 0
		}

		t.speed = t.current.DraggedY
		t.angle = 0
		return
	}

	// forward/backward calculation
	if (t.request.DraggedY > 10 || t.request.DraggedY < -10) &&
		(t.request.DraggedX < 30 && t.request.DraggedX > -30) {

		if t.request.DraggedY != t.current.DraggedY {

			// Calculate a new forward/backward movement value
			// based on the	difference between the request and current
			if t.request.DraggedY > t.current.DraggedY {
				// if the request is further than the current
				diff := t.request.DraggedY - t.current.DraggedY
				if diff > 10 {
					t.current.DraggedY += 10
				} else {
					t.current.DraggedY = t.request.DraggedY
				}
			} else {
				if t.request.DraggedY > 0 {
					// if the request is closer than the current (both are positive)
					diff := t.current.DraggedY - t.request.DraggedY
					if diff > 10 {
						t.current.DraggedY -= 10
					} else {
						t.current.DraggedY = t.request.DraggedY
					}
				} else {
					diff := math.Abs(float64(t.request.DraggedY)) - math.Abs(float64(t.current.DraggedY))
					if diff > 10 {
						t.current.DraggedY -= 10
					} else {
						t.current.DraggedY = t.request.DraggedY
					}
				}
			}
		}
		t.speed = t.current.DraggedY
		t.angle = t.current.DraggedX
		return
	}

	// pure rotation
	if (t.request.DraggedY < 10 && t.request.DraggedY > -10) &&
		(t.request.DraggedX > 30 || t.request.DraggedX < -30) {

		// limiter
		if t.request.DraggedX > 50 {
			t.request.DraggedX = 50
		} else if t.request.DraggedX < -50 {
			t.request.DraggedX = -50
		}

		t.current.DraggedX = 1
		t.current.DraggedY = t.request.DraggedX * -1

		t.speed = t.current.DraggedY
		t.angle = t.current.DraggedX
		return
	}

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
}