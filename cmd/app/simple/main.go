package main

import (
	"log"
	"sync"
	"time"

	"github.com/bus710/tortuga"
	"github.com/bus710/tortuga/cmd/command"
	"github.com/bus710/tortuga/internal/model"
)

// App to hold/connect variables and methods during the lifecycle
type App struct {
	conn tortuga.Connection
}

func main() {

	log.Println("Hello!")

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
		app.conn.Send(command.BaseControlCommand(100, 0))
		time.Sleep(time.Millisecond * 100)
	}

	app.conn.Send(command.BaseControlCommand(0, 0))

	app.conn.Stop()

	waitInstance.Wait()

	log.Println("Bye!")
}

func (app *App) handler(feedback model.Feedback) {
	// log.Println()
	// log.Printf("Available Contents: %32b", feedback.AvailableContent)
	// log.Println("0. Time of processing: ", feedback.TimeStamp)
	log.Println("1. Basic Sensor Data: ", feedback.BasicSensorData)
	// log.Println("3. Docking IR: ", feedback.DockingIR)
	// log.Println("4. Inertial Sensor: ", feedback.InertialSensor)
	// log.Println("5. Cliff: ", feedback.Cliff)
	// log.Println("6. Current: ", feedback.Current)
	// log.Println("10. HW Ver: ", feedback.HardwareVersion)
	// log.Println("11. FW Ver: ", feedback.FirmwareVersion)
	// log.Println("13. Gyro: ", feedback.Gyro.FollowedDataLength/3, feedback.RawGyroDataArray)
	// log.Println("16. GPInput: ", feedback.GeneralPurposeInput)
	// log.Println("19. UDID: ", feedback.UniqueDeviceIdentifier)
	// log.Println("21. Controller Info: ", feedback.ControllerInfo)
}
