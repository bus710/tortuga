package main

import (
	"log"
	"math"
	"sync"
	"time"

	"github.com/bus710/tortuga"
	"github.com/bus710/tortuga/cmd/command"
	"github.com/bus710/tortuga/internal/model"
)

// App ...
type App struct {
	conn tortuga.Connection
}

func main() {
	log.Println("Hello")

	waitInstance := sync.WaitGroup{}

	app := App{}
	app.conn = tortuga.Connection{}
	err := app.conn.Init(&waitInstance, app.handler, "ttyUSB0")
	if err != nil {
		log.Fatal(err)
	}

	waitInstance.Add(1)
	go app.conn.Run()

	for {
		for i := 0.0; i < 6.28; i += 0.04 {
			j := int16(math.Sin(float64(i)) * 120)
			app.conn.Send(command.BaseControlCommand(j, 0))
			time.Sleep(time.Millisecond * 100)
		}
	}

	app.conn.Stop()

	waitInstance.Wait()

	log.Println("Bye!")

}

func (app *App) handler(fdb model.Feedback) {}
