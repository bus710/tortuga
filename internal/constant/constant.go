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
