package tortuga

import (
	"log"
	"sync"
	"testing"
	"time"

	"github.com/bus710/tortuga"
	"github.com/bus710/tortuga/intenal/model"
)

type testHelper struct {
	conn tortuga.Connection
}

func Test(t *testing.T) {

	log.Println("Hello!")

	waitInstance := sync.WaitGroup{}

	tHelper := testHelper{}
	tHelper.conn = tortuga.Connection{}
	err := tHelper.conn.Init(&waitInstance, tHelper.handler, "ttyUSB0")
	if err != nil {
		log.Fatal(err)
	}

	waitInstance.Add(1)
	go tHelper.conn.Run()

	time.Sleep(time.Second * 5)
	tHelper.conn.Stop()

	waitInstance.Wait()

	log.Println("Bye!")
}

func (t *testHelper) handler(packet model.Packet) {
	log.Println("hi")
}
