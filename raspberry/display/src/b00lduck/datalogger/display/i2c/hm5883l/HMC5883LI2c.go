package hmc5883l
import (
	"b00lduck/datalogger/display/i2c"
	"unsafe"
)

const (
	ADDR          		= 0x1E
	REGISTER_A			= 0x00
	REGISTER_A_VALUE    = 0x10
	REGISTER_B			= 0x01
	REGISTER_B_VALUE	= 0x20
	REGISTER_MODE		= 0x02
	REGISTER_MODE_VALUE	= 0x00
	REGISTER_X			= 0x03
	REGISTER_Y			= 0x07
	REGISTER_Z			= 0x05
)

type HMC5883LI2C struct {
	bus *i2c.I2CBus
}

type Vector3 struct {
	X,Y,Z int16
}

func CreateHMC5883LI2c(busId byte) (dev *HMC5883LI2C, err error) {

	devI2c := new(HMC5883LI2C)

	devI2c.bus, err = i2c.Create(busId)
	if err != nil {
		return nil, err
	}

	devI2c.config()
	if err != nil {
		return nil, err
	}

	return devI2c, nil
}

func (dev *HMC5883LI2C) config() (err error) {
	err = dev.Write(REGISTER_A, REGISTER_A_VALUE)
	if err != nil {
		return
	}

	err = dev.Write(REGISTER_B, REGISTER_B_VALUE)
	if err != nil {
		return
	}

	err = dev.Write(REGISTER_MODE, REGISTER_MODE_VALUE)
	if err != nil {
		return
	}
	return
}

func (dev *HMC5883LI2C) ReadVector() (vector *Vector3, err error) {
	vector = new(Vector3)
	vector.X, err = dev.Read16(REGISTER_X)
	if err != nil {
		return
	}
	vector.Y, err = dev.Read16(REGISTER_Y)
	if err != nil {
		return
	}
	vector.Z, err = dev.Read16(REGISTER_Z)
	if err != nil {
		return
	}
	return
}

func (dev *HMC5883LI2C) Read16(reg byte) (int16, error) {

	msb, err := dev.Read(reg)
	if err != nil {
		return 0, err
	}

	lsb, err := dev.Read(reg + 1)
	if err != nil {
		return 0, err
	}

	var x uint16 = (uint16(msb) << 8) + uint16(lsb)

	pxs := (*int16)(unsafe.Pointer(&x))

	return int16(*pxs), nil
}

func (dev *HMC5883LI2C) Read(reg byte) (ret uint8, err error) {
	var bytes []byte
	bytes, err = dev.bus.ReadByteBlock(ADDR, reg, 1)
	ret = uint8(bytes[0])
	return
}

func (dev *HMC5883LI2C) Write(reg byte, value byte) (err error) {
	err = dev.bus.WriteByte(ADDR, reg, value)
	if err != nil {
		return
	}
	return
}