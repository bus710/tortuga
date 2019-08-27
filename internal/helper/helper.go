package helper

import (
	"github.com/bus710/tortuga/internal/model"
)

// Serialize receives a command struct and makes it a byte slice to be used by the serial write
func Serialize(cmd model.Command) (data []byte) {
	data = make([]byte, 0)

	// Headers
	data = append(data, 0xAA)
	data = append(data, 0x55)

	// Total Length
	data = append(data, cmd.Length)

	// ID
	data = append(data, cmd.ID)

	// Size of the following params
	data = append(data, cmd.Size)

	// Playload and CRC
	cmd.CRC = 0
	for i := 0; i < int(cmd.Size); i++ {
		data = append(data, cmd.Payload[i])
		cmd.CRC = cmd.CRC ^ cmd.Payload[i]
	}
	data = append(data, cmd.CRC)

	return data
}

// MergeResidue ...
func MergeResidue(residue []byte, numRead uint16, buf []byte) (newNumRead uint16, newBuf []byte) {
	newNumRead += uint16(len(residue))
	newBuf = append(residue, buf...)

	return newNumRead, newBuf
}

// SearchHeader ...
func SearchHeader(numRead uint16, buf []byte) (pLoc []uint16) {
	pLoc = make([]uint16, 0)
	pLoc = append(pLoc, 0) // preamble consists of 2 digits

	for i := uint16(0); i < numRead; i++ {
		if buf[i] == 0xaa && buf[i+1] == 0x55 {
			pLoc = append(pLoc, i)
			i += 80
		}
	}

	pLoc = append(pLoc, uint16(numRead+1)) // To make the tail to be scan in the next step
	pLoc = append(pLoc, 0)

	return pLoc
}
