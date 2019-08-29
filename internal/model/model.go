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
	LeftPWM          byte
	RightPWM         byte
	Button           byte
	Charger          byte
	Battery          byte
	OvercurrentFlags byte
}

// DockingIR ...
type DockingIR struct {
	RightSignal   byte
	CentralSignal byte
	LeftSignal    byte
}

// InertialSensor ...
type InertialSensor struct {
	Angle     uint16
	AngleRate uint16
}

// Cliff ...
type Cliff struct {
	RightCliffSensor   uint16
	CentralCliffSensor uint16
	LeftCliffSensor    uint16
}

// Current ...
type Current struct {
	LeftMotor  byte
	RightMotor byte
}

// HardwareVersion ...
type HardwareVersion struct {
	Patch byte
	Minor byte
	Major byte
}

// FirmwareVersion ...
type FirmwareVersion struct {
	Patch byte
	Minor byte
	Major byte
}

// Gyro ...
type Gyro struct {
	FrameID                byte
	RawGyroDataArrayLength byte
	RawGyroDataArray       [3]RawGyroData
}

// RawGyroData ...
type RawGyroData struct {
	X uint16
	Y uint16
	Z uint16
}

// GeneralPurposeInput ...
type GeneralPurposeInput struct {
	DigitalInput   uint16
	AnalogInputCH0 uint16
	AnalogInputCH1 uint16
	AnalogInputCH2 uint16
	AnalogInputCH3 uint16
}

// UniqueDeviceIdentifier ...
type UniqueDeviceIdentifier struct {
	UDID0 uint32
	UDID1 uint32
	UDID2 uint32
}

// ControllerInfo ...
type ControllerInfo struct {
	Type  byte
	PGain uint32
	IGain uint32
	DGain uint32
}
