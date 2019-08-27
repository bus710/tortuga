package tortuga

import (
	"errors"
	"io/ioutil"
	"log"
	"strings"
	"sync"
	"time"

	helper "github.com/bus710/tortuga/internal/helper"
	model "github.com/bus710/tortuga/internal/model"
	"github.com/tarm/serial"
)

// Connection ...
type Connection struct {
	test bool

	wait        *sync.WaitGroup
	devName     string
	chanStop    chan bool
	chanCommand chan model.Command

	serialport   *serial.Port
	serialconfig *serial.Config

	numRead int
	buf     []byte
	pLoc    []uint16 // Pleamble Location
	residue []byte   // Used if there is a leftover bytes after parsing

	handler func()
}

// Init this checks available ports and opens one if exists
func (c *Connection) Init(
	wait *sync.WaitGroup, devName string, handler func(), test bool) (err error) {

	c.test = test
	c.wait = wait
	c.chanStop = make(chan bool, 1)
	c.chanCommand = make(chan model.Command, 1)

	// 1. Check if the given name has the pattern expected (/dev/ttyUSB0)
	if !strings.Contains(devName, "tty") {
		err := errors.New("the given device name doens't point a tty device file")
		return err
	}

	// 2. Accept the given name and store it in the struct
	c.devName = devName

	// 3. Check if there ia a device file as the given name
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
		err := errors.New("cannot find a device file as the given name")
		return err
	}

	// 4. Config and open a serial port with the given name
	c.serialconfig = &serial.Config{
		Name: "/dev/" + c.devName,
		Baud: 115200,
	}

	c.serialport, err = serial.OpenPort(c.serialconfig)
	if err != nil {
		c.serialport = nil
		log.Println("issue with the serialport")
		return err
	}

	c.handler = handler

	return nil
}

// Run ...
func (c *Connection) Run() {

	if c.serialport == nil {
		c.wait.Done()
		log.Println("issue with the serialport")
		return
	}

	defer c.serialport.Close()

	ticker := time.NewTicker(100 * time.Millisecond).C
	tickerCount := int(0)

loopRun:
	for {
		select {
		// This case periodically runs the read routine
		case <-ticker:
			if c.test {
				tickerCount++
				if tickerCount > 10 {
					c.wait.Done()
					return
				}
			}

			c.handler()
		// This case receives the command struct from the app
		case command := <-c.chanCommand:
			data, err := helper.Serialize(command)
			if err != nil {
				return
			}
			if data[0] == 1 {
			}
		// This case receives a stop signal
		case <-c.chanStop:
			break loopRun
		}
	}
	c.wait.Done()
}
