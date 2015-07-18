package i2c

import (
	"os"
	"b00lduck/datalogger/display/tools"
)

const (
	I2C_SLAVE = 0x0703
	I2C_SMBUS = 0x0720
)

type I2C struct {
	file *os.File
}

func New(device string, address uint8) (*I2C, error) {
	f, err := os.OpenFile(device, os.O_RDWR, 0600)
	if err != nil {
		return nil, err
	}

	err = tools.Ioctl(f.Fd(), I2C_SLAVE, uintptr(address))
	if err != nil {
		return nil, err
	}

	return &I2C{f}, nil
}

func (i2c *I2C) SendCommand(command, value uint8) {
	buf := make ([]uint8, 2)

	buf[0] = command
	buf[1] = value

	i2c.Write(buf)
}

func (i2c *I2C) Write(buf []byte) (int, error) {
	return i2c.file.Write(buf)
}

func (i2c *I2C) WriteByte(b byte) (int, error) {
	var buf [1]byte
	buf[0] = b
	return i2c.file.Write(buf[:])
}

func (i2c *I2C) Read(p []byte) (int, error) {
	return i2c.file.Read(p)
}

func (i2c *I2C) Close() error {
	return i2c.file.Close()
}
