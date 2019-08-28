package tortuga

import (
	"errors"

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
		return errors.New("written length is not matched")
	}
	return nil
}

// readPort is written to start the marshaling
func (c *Connection) readPort() (err error) {

	c.buf = make([]byte, 8192)
	_, err = c.serialport.Read(c.buf)
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
	c.numRead += uint16(len(c.residue))
	c.buf = append(c.residue, c.buf...)

	return nil
}

func (c *Connection) searchHeader() (err error) {

	/* Preamble Location */
	c.pLoc = make([]uint16, 0)
	c.pLoc = append(c.pLoc, 0)

	for i := uint16(0); i < uint16(c.numRead); i++ {
		if c.buf[i] == 0xaa && c.buf[i+1] == 0x55 {
			c.pLoc = append(c.pLoc, i)
			i += 80
		}
	}

	c.pLoc = append(c.pLoc, uint16(c.numRead+1))
	c.pLoc = append(c.pLoc, 0)

	return nil
}

func (c *Connection) dividePacket() (err error) {

	/* */
	for i, start := range c.pLoc {
		if i+1 == len(c.pLoc) {
			break
		}

		end := c.pLoc[i+1]
		if end != 0 {
			if c.checkCRC(start, end) {
				// log.Printf("%d, %d, %d - %x \n", i, start, end, t.buf[start:end])
				c.residue = make([]byte, 0)
			} else {
				// log.Println("CRC error")
				c.residue = c.buf[start:end]
			}
		}
	}

	return nil
}

func (c *Connection) checkCRC(start, end uint16) bool {

	crc := c.buf[start+2]
	for i := start + 3; i < end-1; i++ {
		crc = crc ^ c.buf[i]
	}

	if c.buf[end-1] == crc {
		return true
	}

	return false
}
