package main

import (
	"log"
	"sync"
)

// App ...
type App struct {
	tortugaInstance Tortuga
	serverInstance  webServer
	signalInstance  termSignal
	waitInstance    sync.WaitGroup
}

func main() {
	log.Println("Hello")

	app := App{}

	app.signalInstance = termSignal{}
	app.tortugaInstance = Tortuga{}
	app.serverInstance = webServer{}

	app.tortugaInstance.init(&app)
	app.signalInstance.init(&app)
	app.serverInstance.init(&app)

	app.waitInstance.Add(1)
	go app.tortugaInstance.run()
	app.waitInstance.Add(1)
	go app.signalInstance.run()
	app.waitInstance.Add(1)
	go app.serverInstance.run()

	app.waitInstance.Wait()

	log.Println("Bye!")
}
