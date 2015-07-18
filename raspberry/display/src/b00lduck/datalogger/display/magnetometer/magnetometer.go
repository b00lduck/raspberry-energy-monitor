package magnetometer

import (
	"b00lduck/datalogger/display/i2c"
	"fmt"
)

type MagnetoEvent struct {
	X int
	Y int
	Z int
}

type Magnetometer struct {
	I2C* i2c.I2C
}

func (f *Magnetometer) Open(device string) {
	i2c, err := i2c.New("/dev/i2c-1", 0x1E)
	if err == nil {
		f.I2C = i2c
	} else {
		f.I2C = nil
	}
}

func (f *Magnetometer) Close() {
	f.I2C.Close()
}

func (f *Magnetometer) Read() error {

	if f.I2C == nil {
		fmt.Println("I2C device not opened, will use dummy data")
	} else {
		fmt.Println("sending command")
		f.I2C.SendCommand(2, 1)
	}

	return nil

}

