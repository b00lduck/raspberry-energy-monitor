package i2c

const (
	ADDR          = 0x1E
	REGISTER_XOUT = 0x06
	REGISTER_YOUT = 0x07
	REGISTER_ZOUT = 0x08
)

type HMC5883L struct {
	bus *I2CBus
}

func New() (dev *HMC5883L, err error) {
	dev = new(HMC5883L)
	dev.bus, err = Bus(1)

	err = dev.Write(0, 0x10)
	err = dev.Write(1, 0x20)
	err = dev.Write(2, 0x00)

	return
}

func (bp *HMC5883L) Read(reg byte) (int8) {
	var bytes []byte
	bytes, _ = bp.bus.ReadByteBlock(ADDR, reg, 1)
	return int8(bytes[0])
}

func (bp *HMC5883L) Write(reg byte, value int8) (err error) {
	err = bp.bus.WriteByte(0x1e, 2, 1)
	if err != nil {
		return
	}
	return
}