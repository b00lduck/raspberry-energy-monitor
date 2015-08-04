package tools
import "fmt"

func ErrorCheck(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}