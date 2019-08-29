// http://yujinrobot.github.io/kobuki/enAppendixProtocolSpecification.html

package constant

// Command ID Constants
const (
	IDBaseControl          = 1
	IDSound                = 3
	IDSoundSequence        = 4
	IDRequestExtra         = 9
	IDGeneralPurposeOutput = 12
	IDSetControllerGain    = 13
	IDGetControllerGain    = 14
)

// Command Total Length Constants
// total length = id + size + payload + crc
const (
	LenBaseControl          = 7
	LenSound                = 6
	LenSoundSequence        = 4
	LenRequestExtra         = 5
	LenGeneralPurposeOutput = 5
	LenSetControllerGain    = 16
	LenGetControllerGain    = 4
)

// Size of the sub payload
const (
	SizeBaseControl          = 4
	SizeSound                = 3
	SizeSoundSequence        = 1
	SizeRequestExtra         = 2
	SizeGeneralPurposeOutput = 2
	SizeSetControllerGain    = 13
	SizeGetControllerGain    = 1
)

// Feedback
const (
	IDBasicSensorData        = 1
	IDDockingIR              = 3
	IDInertialSensor         = 4
	IDCliff                  = 5
	IDCurrent                = 6
	IDHardwareVersion        = 10
	IDFirmwareVersion        = 11
	IDRawDataOf3AxisGyro     = 13
	IDGeneralPurposeInput    = 16
	IDUniqueDeviceIdentifier = 19
	IDControllerInfo         = 21
)

// Size of data field
// Gyro sensor value can be 14 or 20
const (
	SizeBasicSensorData        = 15
	SizeDockingIR              = 3
	SizeInertialSensor         = 7
	SizeCliff                  = 6
	SizeCurrent                = 2
	SizeHardwareVersion        = 4
	SizeFirmwareVersion        = 4
	SizeRawDataOf3AxisGyroA    = 14
	SizeRawDataOf3AxisGyroB    = 20
	SizeGeneralPurposeInput    = 16
	SizeUniqueDeviceIdentifier = 12
	SizeControllerInfo         = 21
)

/* Feedback Example
- Row data:
	aa554d010f90f1000000ed2b58470d0d00129f000303000000
	04073c1dfcff0000000506f506f907900606020101
	0d0e8106a1ff0800c8ff90ff0300cfff
	10100f00dc0fe00fe00fe00fef0f00000000a100

- Preambles: aa55
- Total length: 4d
- Basic Sensor Data: 01 0f 90f1 00 00 00 ed2b 5847 0d 0d 00 12 9f 00
- Docking IR: 03 03 00 00 00
- Inertial Sensor: 04 07 3c1d fcff 00 00 00
- Cliff: 05 06 f506 f907 9006
- Current: 06 02 0101
- Hardware version: not requested
- Firmware version: not requested
- Gyro: 0d 0e 81 06 a1ff 0800 c8ff 90ff 0300 cfff
- General purpose input: 10 10 0f00 dc0f e00f e00f e00f ef0f 0000 0000
- CRC and an empty byte: a1 00 */
