package command

// func (t *tortuga) SendBaseControlCommand(
// 	speed, radius int16) (
// 	cmd Command) {

// 	cmd.length = lenBaseControl
// 	cmd.id = idBaseControl
// 	cmd.size = sizeBaseControl

// 	cmd.payload[0] = byte(speed & 0xff)
// 	cmd.payload[1] = byte((speed >> 8) & 0xff)
// 	cmd.payload[2] = byte(radius & 0xff)
// 	cmd.payload[3] = byte((radius >> 8) & 0xff)

// 	return cmd
// }

// func (t *tortuga) SendSoundCommand(
// 	f, a, duration uint8) (
// 	cmd Command) {

// 	cmd.length = lenSound
// 	cmd.id = idSound
// 	cmd.size = sizeSound

// 	tmp := uint16(1 / (f * a))

// 	cmd.payload[0] = byte(tmp & 0xff)
// 	cmd.payload[1] = byte((tmp >> 8) & 0xff)
// 	cmd.payload[2] = byte(duration)

// 	return cmd
// }

// func (t *tortuga) SendSoundSequenceCommand(
// 	sequence uint8) (
// 	cmd Command) {

// 	cmd.length = lenSoundSequence
// 	cmd.id = idSoundSequence
// 	cmd.size = sizeSoundSequence

// 	cmd.payload[0] = byte(sequence)

// 	return cmd
// }

// func (t *tortuga) SendRequestExtraCommand(
// 	hwVer, fwVer, udid bool) (
// 	cmd Command) {

// 	cmd.length = lenRequestExtra
// 	cmd.id = idRequestExtra
// 	cmd.size = sizeRequestExtra

// 	tmp := byte(0x00)

// 	if hwVer {
// 		tmp |= 0x01
// 	}
// 	if fwVer {
// 		tmp |= 0x02
// 	}
// 	if udid {
// 		tmp |= 0x08
// 	}

// 	cmd.payload[0] = tmp

// 	return cmd
// }

// func (t *tortuga) SendGeneralPurposeOutputCommand(
// 	digitalOutput0,
// 	digitalOutput1,
// 	digitalOutput2,
// 	digitalOutpu3 bool,
// 	power3v3,
// 	power5v0,
// 	power12va,
// 	power12vb bool,
// 	redLed1,
// 	greenLed1,
// 	redLed2,
// 	greenLed2 bool) (cmd Command) {

// 	if digitalOutput {

// 	}
// 	return cmd
// }
