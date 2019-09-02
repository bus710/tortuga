package main

import (
	"log"
	"math"
	"time"

	"github.com/bus710/tortuga"
	"github.com/bus710/tortuga/cmd/command"
	"github.com/bus710/tortuga/cmd/model"
)

// Tortuga ...
type Tortuga struct {
	app      *App
	conn     tortuga.Connection
	chanStop chan bool
}

func (t *Tortuga) init(app *App) {

	t.app = app
	t.conn = tortuga.Connection{}
	t.chanStop = make(chan bool, 1)

	err := t.conn.Init(&app.waitInstance, t.handler, "ttyUSB0")
	if err != nil {
		log.Fatal(err)
	}
}

func (t *Tortuga) run() {

	go t.conn.Run()

	ticker := time.NewTicker(100 * time.Millisecond).C
	cycle := 0.0
run:
	for {
		select {
		case <-ticker:
			cycle += 0.04
			if cycle > 6.28 {
				cycle = 0.0
			}
			swing := int16(math.Sin(float64(cycle)) * 120)
			t.conn.Send(command.BaseControlCommand(swing, 0))

		case <-t.chanStop:
			t.conn.Stop()
			time.Sleep(time.Second)
			break run
		}
	}
	t.app.waitInstance.Done()
}

func (t *Tortuga) handler(fdb model.Feedback) {
}
