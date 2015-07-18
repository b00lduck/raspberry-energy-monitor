package hmc5883l

type HMC5883LMock struct {

}

func CreateHMC5883LMock() (*HMC5883LMock) {
	dev := new(HMC5883LMock)
	return dev
}

func (dev *HMC5883LMock) ReadVector() (vector *Vector3, err error) {
	vector = new(Vector3)
	vector.X = 0
	vector.Y = 32768
	vector.Z = 65535
	err = nil
	return
}