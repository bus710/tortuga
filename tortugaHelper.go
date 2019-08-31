package tortuga

import (
	"errors"
	"time"

	constant "github.com/bus710/tortuga/internal/constant"
	model "github.com/bus710/tortuga/internal/model"
)

// writePort is written to protect the write function
func (c *Connection) writePort(data []byte) (err error) {

	if len(data) > 64 {
		return errors.New("too long data")
	}

	writtenLen, err := c.serialport.Write(data)
	if err != nil {
		return err
	}

	if writtenLen != len(data) {
		return errors.New("written length is not matched to the data size")
	}
	return nil
}

// readPort is written to start the marshaling
func (c *Connection) readPort() (err error) {

	c.buf = make([]byte, 8192)
	c.numRead, err = c.serialport.Read(c.buf)
	if err != nil {
		return err
	}

	return nil
}

func (c *Connection) serialize(cmd model.Command) (data []byte) {
	data = make([]byte, 0)

	// headers
	data = append(data, 0xAA)
	data = append(data, 0x55)

	// leangth
	data = append(data, cmd.Length)

	// base control
	data = append(data, cmd.ID)
	data = append(data, cmd.Size)

	// CRC calculation
	cmd.CRC = 0
	for i := 0; i < int(cmd.Size); i++ {
		data = append(data, cmd.Payload[i])
		cmd.CRC = cmd.CRC ^ cmd.Payload[i]
	}
	data = append(data, cmd.CRC)

	return data
}

func (c *Connection) mergeResidue() (err error) {
	if c.residue != nil {
		// c.buf = append(c.residue, c.buf...)
		c.buf = append(c.residue, c.buf[0:c.numRead]...)
		c.numRead += len(c.residue)
		c.residue = nil
	}
	return nil
}

func (c *Connection) searchHeader() (err error) {

	/* Preamble Location */
	c.pLoc = make([]int, 0)
	c.pLoc = append(c.pLoc, 0)

	for i := 1; i < c.numRead; i++ {
		if (c.buf[i-1] == 0xaa) && (c.buf[i] == 0x55) {
			c.pLoc = append(c.pLoc, i-1)
		}
	}

	return nil
}

func (c *Connection) dividePacket() (err error) {

	if len(c.pLoc) < 2 {
		return
	}

	for i, start := range c.pLoc {

		// If reach to the end of the slice,
		// quit the loop to avoid the out of range error
		if i+1 == len(c.pLoc) {
			c.residue = c.buf[start:]
			break
		} else {
			c.residue = nil
		}

		// end must be the byte right before of the next start byte
		end := c.pLoc[i+1]

		if (end - start) > 70 {
			if c.checkCRC(start, end) {
				// log.Println(c.pLoc, i, start, end)
				c.formatFeedback(start, end)
				// c.residue = make([]byte, 0)
			} else {
				// Should we ignore this bytes because the CRC is not correct? We'll see...
				// c.residue = c.buf[start:end]
			}
		}
	}

	return nil
}

func (c *Connection) checkCRC(start, end int) bool {

	crc := c.buf[start+2]
	for i := start + 3; i < end-1; i++ {
		crc = crc ^ c.buf[i]
	}

	if c.buf[end-1] == crc {
		return true
	}

	return false
}

// Please check the feedback sample attached at the bottom of this file.
func (c *Connection) formatFeedback(start, end int) {

	// Row data
	// log.Printf(">> %d, %d - %x \n", start, end, c.buf[start:end])

	tmp := c.buf[start:end] // To use the part of the entire buffer
	totalLength := tmp[2]
	index := int(3) // To ignore the preambles and the total length

	fdb := model.Feedback{}
	fdb.AvailableContent = (1 << constant.IDTimeStamp)
	fdb.TimeStamp = time.Now()

	exitCount := 0

	for {

		exitCount++
		if index >= int(totalLength) || exitCount > 20 {
			break
		}

		switch tmp[index] {
		case constant.IDBasicSensorData:
			fdb.AvailableContent |= (1 << constant.IDBasicSensorData)
			fdb.BasicSensorData.TimeStamp = uint16(tmp[index+2])
			fdb.BasicSensorData.TimeStamp |= (uint16(tmp[index+3]) << 8)
			fdb.BasicSensorData.Bumper = tmp[index+4]
			fdb.BasicSensorData.WheelDrop = tmp[index+5]
			fdb.BasicSensorData.Cliff = tmp[index+6]
			fdb.BasicSensorData.LeftEncoder = uint16(tmp[index+7])
			fdb.BasicSensorData.LeftEncoder |= (uint16(tmp[index+8]) << 8)
			fdb.BasicSensorData.RightEncoder = uint16(tmp[index+9])
			fdb.BasicSensorData.RightEncoder |= (uint16(tmp[index+10]) << 8)
			fdb.BasicSensorData.LeftPWM = tmp[index+11]
			fdb.BasicSensorData.RightPWM = tmp[index+12]
			fdb.BasicSensorData.Button = tmp[index+13]
			fdb.BasicSensorData.Charger = tmp[index+14]
			fdb.BasicSensorData.Battery = tmp[index+15]
			fdb.BasicSensorData.OvercurrentFlags = tmp[index+16]
			index = index + constant.SizeBasicSensorData + 2
		case constant.IDDockingIR:
			fdb.AvailableContent |= (1 << constant.IDDockingIR)
			fdb.DockingIR.RightSignal = tmp[index+2]
			fdb.DockingIR.CentralSignal = tmp[index+3]
			fdb.DockingIR.LeftSignal = tmp[index+4]
			index = index + constant.SizeDockingIR + 2
		case constant.IDInertialSensor:
			fdb.AvailableContent |= (1 << constant.IDInertialSensor)
			fdb.InertialSensor.Angle = uint16(tmp[index+2])
			fdb.InertialSensor.Angle |= (uint16(tmp[index+3]) << 8)
			fdb.InertialSensor.AngleRate = uint16(tmp[index+4])
			fdb.InertialSensor.AngleRate |= (uint16(tmp[index+5]) << 8)
			index = index + constant.SizeInertialSensor + 2
		case constant.IDCliff:
			fdb.AvailableContent |= (1 << constant.IDCliff)
			fdb.Cliff.RightCliffSensor = uint16(tmp[index+2])
			fdb.Cliff.RightCliffSensor |= (uint16(tmp[index+3]) << 8)
			fdb.Cliff.CentralCliffSensor = uint16(tmp[index+4])
			fdb.Cliff.CentralCliffSensor |= (uint16(tmp[index+5]) << 8)
			fdb.Cliff.LeftCliffSensor = uint16(tmp[index+6])
			fdb.Cliff.LeftCliffSensor |= (uint16(tmp[index+7]) << 8)
			index = index + constant.SizeCliff + 2
		case constant.IDCurrent:
			fdb.AvailableContent |= (1 << constant.IDCurrent)
			fdb.Current.LeftMotor = tmp[index+2]
			fdb.Current.RightMotor = tmp[index+3]
			index = index + constant.SizeCurrent + 2
		case constant.IDHardwareVersion:
			fdb.AvailableContent |= (1 << constant.IDHardwareVersion)
			fdb.HardwareVersion.Patch = tmp[index+2]
			fdb.HardwareVersion.Minor = tmp[index+3]
			fdb.HardwareVersion.Major = tmp[index+4]
			index = index + constant.SizeHardwareVersion + 2
		case constant.IDFirmwareVersion:
			fdb.AvailableContent |= (1 << constant.IDFirmwareVersion)
			fdb.FirmwareVersion.Patch = tmp[index+2]
			fdb.FirmwareVersion.Minor = tmp[index+3]
			fdb.FirmwareVersion.Major = tmp[index+4]
			index = index + constant.SizeFirmwareVersion + 2
		case constant.IDRawDataOf3AxisGyro:
			fdb.AvailableContent |= (1 << constant.IDRawDataOf3AxisGyro)
			fdb.Gyro.FrameID = tmp[index+2]
			fdb.Gyro.FollowedDataLength = tmp[index+3]

			fdb.Gyro.RawGyroDataArray[0].X = uint16(tmp[index+4])
			fdb.Gyro.RawGyroDataArray[0].X |= (uint16(tmp[index+5]) << 8)
			fdb.Gyro.RawGyroDataArray[0].Y = uint16(tmp[index+6])
			fdb.Gyro.RawGyroDataArray[0].Y |= (uint16(tmp[index+7]) << 8)
			fdb.Gyro.RawGyroDataArray[0].Z = uint16(tmp[index+8])
			fdb.Gyro.RawGyroDataArray[0].Z |= (uint16(tmp[index+9]) << 8)
			fdb.Gyro.RawGyroDataArray[1].X = uint16(tmp[index+10])
			fdb.Gyro.RawGyroDataArray[1].X |= (uint16(tmp[index+11]) << 8)
			fdb.Gyro.RawGyroDataArray[1].Y = uint16(tmp[index+12])
			fdb.Gyro.RawGyroDataArray[1].Y |= (uint16(tmp[index+13]) << 8)
			fdb.Gyro.RawGyroDataArray[1].Z = uint16(tmp[index+14])
			fdb.Gyro.RawGyroDataArray[1].Z |= (uint16(tmp[index+15]) << 8)

			if tmp[index+1] == constant.SizeRawDataOf3AxisGyroA {
				index = index + constant.SizeRawDataOf3AxisGyroA + 2
			} else if tmp[index+1] == constant.SizeRawDataOf3AxisGyroB {
				fdb.Gyro.RawGyroDataArray[2].X = uint16(tmp[index+16])
				fdb.Gyro.RawGyroDataArray[2].X |= (uint16(tmp[index+17]) << 8)
				fdb.Gyro.RawGyroDataArray[2].Y = uint16(tmp[index+18])
				fdb.Gyro.RawGyroDataArray[2].Y |= (uint16(tmp[index+19]) << 8)
				fdb.Gyro.RawGyroDataArray[2].Z = uint16(tmp[index+20])
				fdb.Gyro.RawGyroDataArray[2].Z |= (uint16(tmp[index+21]) << 8)
				index = index + constant.SizeRawDataOf3AxisGyroB + 2
			}

		case constant.IDGeneralPurposeInput:
			fdb.AvailableContent |= (1 << constant.IDGeneralPurposeInput)
			fdb.GeneralPurposeInput.DigitalInput = uint16(tmp[index+2])
			fdb.GeneralPurposeInput.DigitalInput |= (uint16(tmp[index+3]) << 8)
			fdb.GeneralPurposeInput.AnalogInputCH0 = uint16(tmp[index+4])
			fdb.GeneralPurposeInput.AnalogInputCH0 |= (uint16(tmp[index+5]) << 8)
			fdb.GeneralPurposeInput.AnalogInputCH1 = uint16(tmp[index+6])
			fdb.GeneralPurposeInput.AnalogInputCH1 |= (uint16(tmp[index+7]) << 8)
			fdb.GeneralPurposeInput.AnalogInputCH2 = uint16(tmp[index+8])
			fdb.GeneralPurposeInput.AnalogInputCH2 |= (uint16(tmp[index+9]) << 8)
			fdb.GeneralPurposeInput.AnalogInputCH3 = uint16(tmp[index+10])
			fdb.GeneralPurposeInput.AnalogInputCH3 |= (uint16(tmp[index+11]) << 8)
			index = index + constant.SizeGeneralPurposeInput + 2
		case constant.IDUniqueDeviceIdentifier:
			fdb.AvailableContent |= (1 << constant.IDUniqueDeviceIdentifier)
			fdb.UniqueDeviceIdentifier.UDID0 = uint32(tmp[index+2])
			fdb.UniqueDeviceIdentifier.UDID0 |= (uint32(tmp[index+3]) << 8)
			fdb.UniqueDeviceIdentifier.UDID0 |= (uint32(tmp[index+4]) << 16)
			fdb.UniqueDeviceIdentifier.UDID0 |= (uint32(tmp[index+5]) << 24)
			fdb.UniqueDeviceIdentifier.UDID1 = uint32(tmp[index+6])
			fdb.UniqueDeviceIdentifier.UDID1 |= (uint32(tmp[index+7]) << 8)
			fdb.UniqueDeviceIdentifier.UDID1 |= (uint32(tmp[index+8]) << 16)
			fdb.UniqueDeviceIdentifier.UDID1 |= (uint32(tmp[index+9]) << 24)
			fdb.UniqueDeviceIdentifier.UDID2 = uint32(tmp[index+10])
			fdb.UniqueDeviceIdentifier.UDID2 |= (uint32(tmp[index+11]) << 8)
			fdb.UniqueDeviceIdentifier.UDID2 |= (uint32(tmp[index+12]) << 16)
			fdb.UniqueDeviceIdentifier.UDID2 |= (uint32(tmp[index+13]) << 24)
			index = index + constant.SizeUniqueDeviceIdentifier + 2
		case constant.IDControllerInfo:
			fdb.AvailableContent |= (1 << constant.IDControllerInfo)
			fdb.ControllerInfo.PGain = uint32(tmp[index+2])
			fdb.ControllerInfo.PGain = (uint32(tmp[index+3]) << 8)
			fdb.ControllerInfo.PGain = (uint32(tmp[index+4]) << 16)
			fdb.ControllerInfo.PGain = (uint32(tmp[index+5]) << 24)
			fdb.ControllerInfo.IGain = uint32(tmp[index+6])
			fdb.ControllerInfo.IGain = (uint32(tmp[index+7]) << 8)
			fdb.ControllerInfo.IGain = (uint32(tmp[index+8]) << 16)
			fdb.ControllerInfo.IGain = (uint32(tmp[index+9]) << 24)
			fdb.ControllerInfo.DGain = uint32(tmp[index+10])
			fdb.ControllerInfo.DGain = (uint32(tmp[index+11]) << 8)
			fdb.ControllerInfo.DGain = (uint32(tmp[index+12]) << 16)
			fdb.ControllerInfo.DGain = (uint32(tmp[index+13]) << 24)
			index = index + constant.SizeControllerInfo + 2
		default:
			// log.Println("Check the raw data...")
			fdb.AvailableContent = 0
		}
	}

	// There is no error found
	// Call the given handler function with the parsed feedback data
	if fdb.AvailableContent != 0 {
		c.handler(fdb)
	}
}

// TestHelper ...
// Don't know why cannot call the unexported function (formatFeedback) from the test
func (c *Connection) TestHelper(start, end int, buf []byte) {
	c.buf = buf
	c.formatFeedback(start, end)
}

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
- CRC and an empty byte: a1 00

Compare the row data and parsed */
