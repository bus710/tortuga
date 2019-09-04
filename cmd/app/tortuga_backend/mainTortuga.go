package main

import (
	"log"
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
	// Compare the request and current to drive smmothly
	// if t.request.OriginalX != 0 && t.request.OriginalX != t.current.OriginalX {
	// 	if t.request.Speed > t.current.Speed {
	// 		diff := t.request.Speed - t.current.Speed
	// 		if diff > 10 {
	// 			t.current.Speed += 10
	// 		} else {
	// 			t.current.Speed += diff
	// 		}
	// 	} else {
	// 		diff := t.current.Speed - t.request.Speed
	// 		if diff > 10 {
	// 			t.current.Speed -= 10
	// 		} else {
	// 			t.current.Speed -= diff
	// 		}
	// 	}
	// }
	// if t.request.Speed == 0 && t.current.Speed != 0 {
	// 	diff := t.current.Speed - 10
	// 	if diff > 10 {
	// 		t.current.Speed -= 10
	// 	} else {
	// 		t.current.Speed -= diff
	// 	}
	// }
}
