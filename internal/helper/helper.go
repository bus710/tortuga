package helper

import (
	"github.com/bus710/tortuga/internal/model"
)

// Serialize ...
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
