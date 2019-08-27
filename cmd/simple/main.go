package main

import (
	"log"
	"sync"

	"github.com/bus710/tortuga"
)

func main() {
	waitInstance := sync.WaitGroup{}
	h := handler

	tConn := tortuga.Connection{}
	tConn.Init(&waitInstance, "/dev/ttyUSB0", h)
}

func handler() {
	log.Println("test")
}
