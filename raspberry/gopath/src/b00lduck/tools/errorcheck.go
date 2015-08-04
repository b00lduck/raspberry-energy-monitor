package tools

func ErrorCheck(e error) {
	if e != nil {
		panic(e)
	}
}