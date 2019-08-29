package model

import "time"

// Command can be used to generate a command for a Kobuki
type Command struct {
	Header  [2]byte
	Length  byte
	ID      byte
	Size    byte
	Payload [15]byte
	CRC     byte
}

// Feedback can be used as a parsed data stream
type Feedback struct {
	AvailableContent uint32
	TimeStamp        time.Time
	BasicSensorData
	DockingIR
	InertialSensor
	Cliff
	Current
	HardwareVersion
	FirmwareVersion
	Gyro
	GeneralPurposeInput
	UniqueDeviceIdentifier
	ControllerInfo
}

// BasicSensorData ...
type BasicSensorData struct {
	TimeStamp        uint16
	Bumper           byte
	WheelDrop        byte
	Cliff            byte
	LeftEncoder      uint16
	RightEncoder     uint16
	LeftPWN          byte
	RightPWM         byte
	Button           byte
	Charger          byte
	Battery          byte
	OvercurrentFlags byte
}

// DockingIR ...
type DockingIR struct{}

// InertialSensor ...
type InertialSensor struct{}

// Cliff ...
type Cliff struct{}

// Current ...
type Current struct{}

// HardwareVersion ...
type HardwareVersion struct{}

// FirmwareVersion ...
type FirmwareVersion struct{}

// Gyro ...
type Gyro struct{}

// GeneralPurposeInput ...
type GeneralPurposeInput struct{}

// UniqueDeviceIdentifier ...
type UniqueDeviceIdentifier struct{}

// ControllerInfo ...
type ControllerInfo struct{}
