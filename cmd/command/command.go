package command

import (
	constant "github.com/bus710/tortuga/cmd/constant"
	model "github.com/bus710/tortuga/cmd/model"
)

// BaseControlCommand moves the body regarding the params
func BaseControlCommand(
	speed, radius int16) (
	cmd model.Command) {

	cmd.Length = constant.LenBaseControl
	cmd.ID = constant.IDBaseControl
	cmd.Size = constant.SizeBaseControl

	// if radius > 100 {
	// 	if speed > 1 {
	// 		speed = speed * (radius + 230/2) / radius
	// 	} else if speed < 1 {
	// 		speed = speed * (radius - 230/2) / radius
	// 	}
	// }

	cmd.Payload[0] = byte(speed & 0xff)
	cmd.Payload[1] = byte((speed >> 8) & 0xff)
	cmd.Payload[2] = byte(radius & 0xff)
	cmd.Payload[3] = byte((radius >> 8) & 0xff)

	return cmd
}

// SoundCommand makes sound in the low level
func SoundCommand(
	f, a, duration uint8) (
	cmd model.Command) {

	cmd.Length = constant.LenSound
	cmd.ID = constant.IDSound
	cmd.Size = constant.SizeSound

	tmp := uint16(1 / (f * a))

	cmd.Payload[0] = byte(tmp & 0xff)
	cmd.Payload[1] = byte((tmp >> 8) & 0xff)
	cmd.Payload[2] = byte(duration)

	return cmd
}

// SoundSequenceCommand makes the robot to sing
func SoundSequenceCommand(
	sequence uint8) (
	cmd model.Command) {

	cmd.Length = constant.LenSoundSequence
	cmd.ID = constant.IDSoundSequence
	cmd.Size = constant.SizeSoundSequence

	cmd.Payload[0] = byte(sequence)

	return cmd
}

// RequestExtraCommand requests the robot's information
func RequestExtraCommand(
	hwVer, fwVer, udid bool) (
	cmd model.Command) {

	cmd.Length = constant.LenRequestExtra
	cmd.ID = constant.IDRequestExtra
	cmd.Size = constant.SizeRequestExtra

	tmp := byte(0x00)

	if hwVer {
		tmp |= 0x01
	}
	if fwVer {
		tmp |= 0x02
	}
	if udid {
		tmp |= 0x08
	}

	cmd.Payload[0] = tmp

	return cmd
}

// GeneralPurposeOutputCommand controls the GPIOs, power outputs, and the LEDs
func GeneralPurposeOutputCommand(
	digitalOutput0,
	digitalOutput1,
	digitalOutput2,
	digitalOutpu3 bool,
	power3v3,
	power5v0,
	power12va,
	power12vb bool,
	redLed1,
	greenLed1,
	redLed2,
	greenLed2 bool) (cmd model.Command) {

	// TODO: implementation
	// if digitalOutput {
	// }
	return cmd
}
