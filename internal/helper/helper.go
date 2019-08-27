package helper

import (
	"github.com/bus710/tortuga/internal/model"
)

// Serialize ...
func Serialize(command model.Command) (data []byte, err error) {
	data = make([]byte, 0)

	// headers
	data = append(data, 0xAA)
	data = append(data, 0x55)

	// length
	data = append(data, command.Length)

	// base control
	data = append(data, command.ID)
	data = append(data, command.Size)

	// CRC calculation
	command.CRC = 0
	for i := 0; i < int(command.Size); i++ {
		data = append(data, command.Payload[i])
		command.CRC = command.CRC ^ command.Payload[i]
	}
	data = append(data, command.CRC)

	return data, nil
}
