package hmc5883l

type HMC5883LMock struct {

}

func CreateHMC5883LMock() (*HMC5883LMock) {
	dev := new(HMC5883LMock)
	return dev
}

func (dev *HMC5883LMock) ReadVector() (vector *Vector3, err error) {
	vector = new(Vector3)
	vector.X = 1000
	vector.Y = 2000
	vector.Z = 5000
	err = nil
	return
}