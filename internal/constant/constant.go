package constant

// Command ID Constants
const (
	idBaseControl          = 1
	idSound                = 3
	idSoundSequence        = 4
	idRequestExtra         = 9
	idGeneralPurposeOutput = 12
	idSetControllerGain    = 13
	idGetControllerGain    = 14
)

// Command Total Length Constants
// total length = id + size + payload + crc
const (
	lenBaseControl          = 7
	lenSound                = 6
	lenSoundSequence        = 4
	lenRequestExtra         = 5
	lenGeneralPurposeOutput = 5
	lenSetControllerGain    = 16
	lenGetControllerGain    = 4
)

// Size of the sub payload
const (
	sizeBaseControl          = 4
	sizeSound                = 3
	sizeSoundSequence        = 1
	sizeRequestExtra         = 2
	sizeGeneralPurposeOutput = 2
	sizeSetControllerGain    = 13
	sizeGetControllerGain    = 1
)
