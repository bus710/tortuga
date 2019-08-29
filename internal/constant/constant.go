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
	IDTimeStamp              = 0
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
