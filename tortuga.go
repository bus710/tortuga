package tortuga

import (
	"errors"
	"io/ioutil"
	"strings"
	"sync"
	"time"

	"github.com/tarm/serial"
	"github.com/bus710/tortuga/internal/helper"
)

// Connection ...
type Connection struct {
	wait        *sync.WaitGroup
	devName     string
	chanStop    chan bool
	chanCommand chan Command

	serialport   *serial.Port
	serialconfig *serial.Config

	numRead int
	buf     []byte
	pLoc    []uint16 // Pleamble Location
	residue []byte   // Used if there is a leftover bytes after parsing
}

// Command can be used to generate a command for a Kobuki
type Command struct {
	header  [2]byte
	length  byte
	id      byte
	size    byte
	payload [15]byte
	crc     byte
}

// Init this checks available ports and opens one if exists
func (c *Connection) Init(
	wait *sync.WaitGroup, devName string) (err error) {

	c.wait = wait
	c.chanStop = make(chan bool, 1)
	c.chanCommand = make(chan Command, 1)

	// 1. Check if the given name has the pattern expected (/dev/ttyUSB0)
	if !strings.Contains(devName, "dev") {
		err := errors.New("the given device name doens't point the /dev directory")
		return err
	}
	if !strings.Contains(devName, "tty") {
		err := errors.New("the given device name doens't point a tty device file")
		return err
	}

	// 2. Accept the given name and store it in the struct
	c.devName = devName

	// 3. Check if there ia a device file as the given name
	devDir, err := ioutil.ReadDir("dev")
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
		return err
	}

	return nil
}

// Run ...
func (c *Connection) Run() (err error) {

	defer c.serialport.Close()

	ticker := time.NewTicker(100 * time.Millisecond).C
	count := int(0)

loopRun:
	for {
		select {
		// This case periodically runs the read routine
		case <-ticker:
			if count == 0 {
			}
		// This case receives the command struct from the app
		case command := <-c.chanCommand:
			data, err := c.serialize(command)
			if err != nil {
				return err
			}
			if data {
			}
			dummy()
		// This case receives a stop signal
		case <-c.chanStop:
			break loopRun
		}
	}

	return nil
}
