package tortuga

import (
	"log"
	"sync"
	"testing"
	"time"

	"github.com/bus710/tortuga"
	"github.com/bus710/tortuga/cmd/command"
	"github.com/bus710/tortuga/internal/model"
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

	for i := 0; i < 5; i++ {
		log.Println(i)
		tHelper.conn.Send(command.BaseControlCommand(50, 0))
		time.Sleep(time.Second * 1)
	}

	tHelper.conn.Send(command.BaseControlCommand(0, 0))
	time.Sleep(time.Second * 2)

	for i := 0; i < 5; i++ {
		log.Println(i)
		tHelper.conn.Send(command.BaseControlCommand(-50, 0))
		time.Sleep(time.Second * 1)
	}

	tHelper.conn.Send(command.BaseControlCommand(0, 0))

	tHelper.conn.Stop()

	waitInstance.Wait()

	log.Println("Bye!")
}

func (t *testHelper) handler(feedback model.Feedback) {
}
