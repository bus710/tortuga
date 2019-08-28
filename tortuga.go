package tortuga

import (
	"errors"
	"io/ioutil"
	"log"
	"strings"
	"sync"
	"time"

	model "github.com/bus710/tortuga/internal/model"
	serial "github.com/tarm/serial"
)

// Connection defines the internal variables
type Connection struct {
	wait    *sync.WaitGroup
	handler func(packet model.Packet)
	devName string

	chanStop    chan bool
	chanCommand chan model.Command

	serialport   *serial.Port
	serialconfig *serial.Config

	numRead int
	buf     []byte
	pLoc    []uint16 // Pleamble Location
	residue []byte   // Used if there is a leftover bytes after parsing

	errCount int
}

// Init this checks available ports and opens one if exists
func (c *Connection) Init(
	wait *sync.WaitGroup, handler func(packet model.Packet), devName string) (err error) {

	c.wait = wait
	c.handler = handler
	c.devName = devName

	c.chanStop = make(chan bool, 1)
	c.chanCommand = make(chan model.Command, 1)

	// 1. Check if the given name has the pattern expected (/dev/ttyUSB0)
	if !strings.Contains(devName, "tty") {
		err := errors.New("the given device name doens't point to a tty device file")
		return err
	}

	// 2. Check if there ia a device file as the given name
	devDir, err := ioutil.ReadDir("/dev")
	if err != nil {
		err := errors.New("cannot read the /dev directory")
		return err
	}

	found := false
	for _, deviceFile := range devDir {
		if strings.Contains(deviceFile.Name(), c.devName) {
			found = true
		}
	}

	if !found {
		err := errors.New("cannot find a device file as the given name - " + c.devName)
		return err
	}

	// 3. Config and open a serial port with the given name
	c.serialconfig = &serial.Config{
		Name: "/dev/" + c.devName,
		Baud: 115200,
	}

	c.serialport, err = serial.OpenPort(c.serialconfig)
	if err != nil {
		c.serialport = nil
		return err
	}

	return nil
}

// Run does these:
// - reads the port opened once 100 ms and passes the data received to parse to be the packet struct
// - serializes a command from the app as a byte slice and writes it to the port opened
// - when the port is disconnected or there is a stop signal (chanStop), stops the loop and exits
func (c *Connection) Run() {

	if c.serialport == nil {
		c.wait.Done()
		return
	}

	defer c.serialport.Close()

	ticker := time.NewTicker(100 * time.Millisecond).C

loopRun:
	for {
		select {
		case <-ticker:
			// Downstream - from app to robot
			err := c.readPort()
			if err != nil {
				c.errCount++
				if c.errCount > 3 {
					// If fail to read the port fails more than 3 times
					break loopRun
				}
			}
			c.mergeResidue()
			c.searchHeader()
			c.dividePacket()

		case command := <-c.chanCommand:
			// Upstream - from robot to app
			data := c.serialize(command)
			err := c.writePort(data)
			if err != nil {
				log.Println(err)
			}

		case <-c.chanStop:
			// Routine cenceler
			break loopRun
		}
	}
	c.wait.Done()
}

// Stop signals to stop the loop
func (c *Connection) Stop() {
	c.chanStop <- true
}

// Send sends command to the robot
func (c *Connection) Send(cmd model.Command) {
	c.chanCommand <- cmd
}
