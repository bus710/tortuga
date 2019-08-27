package main

import (
	"log"
	"sync"

	"github.com/bus710/tortuga"
)

func main() {
	log.Println("Hello!")

	waitInstance := sync.WaitGroup{}

	tConn := tortuga.Connection{}
	err := tConn.Init(&waitInstance, "ttyUSB0", handler, true)
	if err != nil {
		log.Fatal(err)
	}

	waitInstance.Add(1)
	go tConn.Run()

	waitInstance.Wait()

	log.Println("Bye!")
}

func handler() {
	log.Println("test")
}
