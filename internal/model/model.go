package model

// Command can be used to generate a command for a Kobuki
type Command struct {
	Header  [2]byte
	Length  byte
	ID      byte
	Size    byte
	Payload [15]byte
	CRC     byte
}

// Freedback can be used as a parsed data stream
type Feedback struct{}
