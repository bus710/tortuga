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
func MergeResidue(residue []byte, buf []byte) (newBuf []byte) {
	newBuf = append(residue, buf...)

	return newBuf
}

// SearchHeader ...
func SearchHeader(buf []byte) (pLoc []uint16) {
	numRead := uint16(len(buf))
	pLoc = make([]uint16, 0)
	pLoc = append(pLoc, 0) // preamble consists of 2 digits

	for i := uint16(0); i < numRead; i++ {
		if buf[i] == 0xaa && buf[i+1] == 0x55 {
			pLoc = append(pLoc, i)
			i += 80
		}
	}

	pLoc = append(pLoc, uint16(numRead+1)) // To make the tail to be scanned in the next step
	pLoc = append(pLoc, 0)                 // To indicate this is the last cell

	return pLoc
}

// DividePacket ...
func DividePacket(pLoc []uint16, buf []byte, handler func(packet model.Packet)) (residue []byte) {

	residue = make([]byte, 0)

	for i, start := range pLoc {
		if i+1 == len(pLoc) {
			break
		}

		end := pLoc[i+1] // This is actually the location of the fist byte of the next packet
		if end != 0 {
			if CheckCRC(start, end, buf) {
				// log.Printf("%d, %d, %d - %x \n", i, start, end, t.buf[start:end])
				p := model.Packet{}
				handler(p)
				residue = make([]byte, 0)
			} else {
				// log.Println("CRC error")
				residue = buf[start:end]
			}
		}
	}

	return residue
}

// CheckCRC ...
func CheckCRC(start, end uint16, buf []byte) bool {

	crc := buf[start+2]
	for i := start + 3; i < end-1; i++ {
		crc = crc ^ buf[i]
	}

	if buf[end-1] == crc {
		return true
	}

	return false
}
