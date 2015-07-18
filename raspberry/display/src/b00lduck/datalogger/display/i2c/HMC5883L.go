package i2c

const (
	ADDR          		= 0x1E
	REGISTER_A			= 0x00
	REGISTER_A_VALUE    = 0x10
	REGISTER_B			= 0x01
	REGISTER_B_VALUE	= 0x20
	REGISTER_MODE		= 0x02
	REGISTER_MODE_VALUE	= 0x00
)

type HMC5883L struct {
	bus *I2CBus
}

type Vector3 struct {
	X,Y,Z uint16
}

func New() (dev *HMC5883L, err error) {
	dev = new(HMC5883L)
	dev.bus, err = Bus(1)
	err = dev.Write(REGISTER_A, REGISTER_A_VALUE)
	err = dev.Write(REGISTER_B, REGISTER_B_VALUE)
	err = dev.Write(REGISTER_MODE, REGISTER_MODE_VALUE)
	return
}

func (dev *HMC5883L) ReadVector() (*Vector3) {

	return Vector3{
		uint16(dev.Read(3)) << 8 + uint16(dev.Read(4)),
		uint16(dev.Read(7)) << 8 + uint16(dev.Read(8)),
		uint16(dev.Read(5)) << 8 + uint16(dev.Read(6))
	}

}

func (dev *HMC5883L) Read(reg byte) (int8) {
	var bytes []byte
	bytes, _ = dev.bus.ReadByteBlock(ADDR, reg, 1)
	return int8(bytes[0])
}

func (dev *HMC5883L) Write(reg byte, value int8) (err error) {
	err = dev.bus.WriteByte(0x1e, byte, value)
	if err != nil {
		return
	}
	return
}