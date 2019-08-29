package tortuga

import (
	"encoding/hex"
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

func Test_formatFeedback(t *testing.T) {

	log.Println("Hello!")

	waitInstance := sync.WaitGroup{}

	tHelper := testHelper{}
	tHelper.conn = tortuga.Connection{}
	err := tHelper.conn.Init(&waitInstance, tHelper.handler, "ttyUSB0")

	decoded, err := hex.DecodeString("aa554d010f90f1000000ed2b58470d0d00129f00030300000004073c1dfcff0000000506f506f9079006060201010d0e8106a1ff0800c8ff90ff0300cfff10100f00dc0fe00fe00fe00fef0f00000000a100")
	if err != nil {
	}

	tHelper.conn.buf = decoded

	tHelper.conn.formatFeedback(0, len(decoded))

	log.Println("Bye")
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
