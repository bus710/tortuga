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
	tHelper.conn.Init(&waitInstance, tHelper.handler, "ttyUSB0")

	decoded, _ := hex.DecodeString("aa554d010f90f1000000ed2b58470d0d00129f00030300000004073c1dfcff0000000506f506f9079006060201010d0e8106a1ff0800c8ff90ff0300cfff10100f00dc0fe00fe00fe00fef0f00000000a100")
	tHelper.conn.TestHelper(0, uint16(len(decoded)), decoded)

	decoded, _ = hex.DecodeString("aa554d010fa4f1000000f42b60470d0d00129f00030300000004073d1d21000000000506ec06f5079306060201010d0e830691ff0b00ddff9dff2000f5ff10100f00d90fdf0fdb0fdf0fef0f000000008f")
	tHelper.conn.TestHelper(0, uint16(len(decoded)), decoded)

	decoded, _ = hex.DecodeString("aa554d010fb8f1000000fc2b67470d0d00129f00030300000004073e1d4f000000000506e406f4079406060201010d0e8506a5ff34001200a3ff39002b0010100f00dc0fe40fdb0fdd0ff00f00000000e7")
	tHelper.conn.TestHelper(0, uint16(len(decoded)), decoded)
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
	log.Println()
	log.Printf("Available Contents: %x", feedback.AvailableContent)
	log.Println("Basic Sensor Data: ", feedback.BasicSensorData)
	log.Println("Docking IR: ", feedback.DockingIR)
	log.Println("Inertial Sensor: ", feedback.InertialSensor)
	log.Println("Cliff: ", feedback.Cliff)
	log.Println("Current: ", feedback.Current)
	log.Println("HW Ver: ", feedback.HardwareVersion)
	log.Println("FW Ver: ", feedback.FirmwareVersion)
	log.Println("Gyro: ", feedback.Gyro.FollowedDataLength/3, feedback.RawGyroDataArray)
	log.Println("GPInput: ", feedback.GeneralPurposeInput)
	log.Println("UDID: ", feedback.UniqueDeviceIdentifier)
	log.Println("Controller Info: ", feedback.ControllerInfo)
}

/* Expected ourput for "go test -v -run Test_formatFeedback"

=== RUN   Test_formatFeedback
2019/08/29 13:15:39 Hello!
2019/08/29 13:15:39
2019/08/29 13:15:39 Available Contents: 1207b
2019/08/29 13:15:39 Basic Sensor Data:  {61840 0 0 0 11245 18264 13 13 0 18 159 0}
2019/08/29 13:15:39 Docking IR:  {0 0 0}
2019/08/29 13:15:39 Inertial Sensor:  {7484 65532}
2019/08/29 13:15:39 Cliff:  {1781 2041 1680}
2019/08/29 13:15:39 Current:  {1 1}
2019/08/29 13:15:39 HW Ver:  {0 0 0}
2019/08/29 13:15:39 FW Ver:  {0 0 0}
2019/08/29 13:15:39 Gyro:  2 [{65441 8 65480} {65424 3 65487} {0 0 0}]
2019/08/29 13:15:39 GPInput:  {15 4060 4064 4064 4064}
2019/08/29 13:15:39 UIDI:  {0 0 0}
2019/08/29 13:15:39 Controller Info:  {0 0 0 0}
2019/08/29 13:15:39
2019/08/29 13:15:39 Available Contents: 1207b
2019/08/29 13:15:39 Basic Sensor Data:  {61860 0 0 0 11252 18272 13 13 0 18 159 0}
2019/08/29 13:15:39 Docking IR:  {0 0 0}
2019/08/29 13:15:39 Inertial Sensor:  {7485 33}
2019/08/29 13:15:39 Cliff:  {1772 2037 1683}
2019/08/29 13:15:39 Current:  {1 1}
2019/08/29 13:15:39 HW Ver:  {0 0 0}
2019/08/29 13:15:39 FW Ver:  {0 0 0}
2019/08/29 13:15:39 Gyro:  2 [{65425 11 65501} {65437 32 65525} {0 0 0}]
2019/08/29 13:15:39 GPInput:  {15 4057 4063 4059 4063}
2019/08/29 13:15:39 UIDI:  {0 0 0}
2019/08/29 13:15:39 Controller Info:  {0 0 0 0}
2019/08/29 13:15:39
2019/08/29 13:15:39 Available Contents: 1207b
2019/08/29 13:15:39 Basic Sensor Data:  {61880 0 0 0 11260 18279 13 13 0 18 159 0}
2019/08/29 13:15:39 Docking IR:  {0 0 0}
2019/08/29 13:15:39 Inertial Sensor:  {7486 79}
2019/08/29 13:15:39 Cliff:  {1764 2036 1684}
2019/08/29 13:15:39 Current:  {1 1}
2019/08/29 13:15:39 HW Ver:  {0 0 0}
2019/08/29 13:15:39 FW Ver:  {0 0 0}
2019/08/29 13:15:39 Gyro:  2 [{65445 52 18} {65443 57 43} {0 0 0}]
2019/08/29 13:15:39 GPInput:  {15 4060 4068 4059 4061}
2019/08/29 13:15:39 UIDI:  {0 0 0}
2019/08/29 13:15:39 Controller Info:  {0 0 0 0}
2019/08/29 13:15:39 Bye
--- PASS: Test_formatFeedback (0.00s)
PASS
ok      github.com/bus710/tortuga/test  0.003s */
