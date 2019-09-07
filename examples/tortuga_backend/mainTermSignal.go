package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type termSignal struct {
	app      *App
	sigTerm  chan os.Signal
	chanStop chan bool
}

func (ts *termSignal) init(app *App) {
	ts.app = app
	ts.sigTerm = make(chan os.Signal, 1)
}

func (ts *termSignal) run() {

	signal.Notify(ts.sigTerm, syscall.SIGINT, syscall.SIGTERM)

	select {
	case received := <-ts.sigTerm:
		log.Println("Receibed a CTRL+C", received)
		ts.cleanup()

	case <-ts.chanStop:
		log.Println("Received a signal internally")
		ts.cleanup()
	}
}

func (ts *termSignal) cleanup() {
	log.Println("Cleanup - started")
	ts.app.tortugaInstance.chanStop <- true
	time.Sleep(time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*300)
	defer cancel()

	ts.app.serverInstance.instance.Shutdown(ctx)
	ts.app.waitInstance.Done()
	log.Println("Cleanup - done")
}
