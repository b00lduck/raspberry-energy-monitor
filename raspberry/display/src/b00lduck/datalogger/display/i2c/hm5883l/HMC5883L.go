package hmc5883l

type HMC5883L interface {
	ReadVector() (vector *Vector3, err error)
}
