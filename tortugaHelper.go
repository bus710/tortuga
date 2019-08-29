package tortuga

import (
	"errors"
	"log"
	"time"

	constant "github.com/bus710/tortuga/internal/constant"
	model "github.com/bus710/tortuga/internal/model"
)

// writePort is written to protect the write function
func (c *Connection) writePort(data []byte) (err error) {

	if len(data) > 64 {
		return errors.New("too long data")
	}

	writtenLen, err := c.serialport.Write(data)
	if err != nil {
		return err
	}

	if writtenLen != len(data) {
		return errors.New("written length is not matched to the data size")
	}
	return nil
}

// readPort is written to start the marshaling
func (c *Connection) readPort() (err error) {

	c.buf = make([]byte, 8192)
	c.numRead, err = c.serialport.Read(c.buf)
	if err != nil {
		return err
	}

	return nil
}

func (c *Connection) serialize(cmd model.Command) (data []byte) {
	data = make([]byte, 0)

	// headers
	data = append(data, 0xAA)
	data = append(data, 0x55)

	// leangth
	data = append(data, cmd.Length)

	// base control
	data = append(data, cmd.ID)
	data = append(data, cmd.Size)

	// CRC calculation
	cmd.CRC = 0
	for i := 0; i < int(cmd.Size); i++ {
		data = append(data, cmd.Payload[i])
		cmd.CRC = cmd.CRC ^ cmd.Payload[i]
	}
	data = append(data, cmd.CRC)

	return data
}

func (c *Connection) mergeResidue() (err error) {
	c.numRead += len(c.residue)
	c.buf = append(c.residue, c.buf...)

	return nil
}

func (c *Connection) searchHeader() (err error) {

	/* Preamble Location */
	c.pLoc = make([]uint16, 0)
	c.pLoc = append(c.pLoc, 0)

	for i := uint16(0); i < uint16(c.numRead); i++ {
		if i < uint16(len(c.buf)) {
			if c.buf[i] == 0xaa && c.buf[i+1] == 0x55 {
				c.pLoc = append(c.pLoc, i)
				i += 80
			}
		}
	}

	c.pLoc = append(c.pLoc, uint16(c.numRead+1))
	c.pLoc = append(c.pLoc, 0)

	return nil
}

func (c *Connection) dividePacket() (err error) {

	for i, start := range c.pLoc {

		if i+1 == len(c.pLoc) {
			break
		}

		end := c.pLoc[i+1]

		if end != 0 {
			if c.checkCRC(start, end) {
				c.formatFeedback(start, end)
				c.residue = make([]byte, 0)
			} else {
				// Because the leftover bytes don't have a correct CRC byte,
				// it is most likely incomplete packet.
				// To have a complete packet,
				// the leftover should be passed to the next data reading iteration.
				c.residue = c.buf[start:end]
			}
		}
	}

	return nil
}

func (c *Connection) checkCRC(start, end uint16) bool {

	crc := c.buf[start+2]
	for i := start + 3; i < end-1; i++ {
		if uint16(len(c.buf)) < i {
			log.Println(i)
		}
		crc = crc ^ c.buf[i]
	}

	if c.buf[end-1] == crc {
		return true
	}

	return false
}

/* Feedback Example
- Row data:
	aa554d010f90f1000000ed2b58470d0d00129f000303000000
	04073c1dfcff0000000506f506f907900606020101
	0d0e8106a1ff0800c8ff90ff0300cfff
	10100f00dc0fe00fe00fe00fef0f00000000a100

- Preambles: aa55
- Total length: 4d
- Basic Sensor Data: 01 0f 90f1 00 00 00 ed2b 5847 0d 0d 00 12 9f 00
- Docking IR: 03 03 00 00 00
- Inertial Sensor: 04 07 3c1d fcff 00 00 00
- Cliff: 05 06 f506 f907 9006
- Current: 06 02 0101
- Hardware version: not requested
- Firmware version: not requested
- Gyro: 0d 0e 81 06 a1ff 0800 c8ff 90ff 0300 cfff
- General purpose input: 10 10 0f00 dc0f e00f e00f e00f ef0f 0000 0000
- CRC and an empty byte: a1 00

Compare the row data and parsed */

func (c *Connection) formatFeedback(start, end uint16) {

	// Row data
	// log.Printf("%d, %d - %x \n", start, end, c.buf[start:end])

	tmp := c.buf[3:end] // ignore the preambles and the total length
	fdb := model.Feedback{}

	fdb.TimeStamp = time.Now()
	fdb.AvailableContent = (1 << constant.IDTimeStamp)
	index := uint16(0)

	for {
		switch tmp[index] {
		case constant.IDBasicSensorData:
			{
				fdb.AvailableContent = (1 << constant.IDBasicSensorData)
				fdb.BasicSensorData.TimeStamp = uint16(tmp[index+2]) << 8
				fdb.BasicSensorData.TimeStamp |= uint16(tmp[index+1])
			}
		}
	}
}
