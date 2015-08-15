package main

import (
	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi"
	"fmt"
	"time"
)


func main() {

	embd.InitGPIO()
	defer embd.CloseGPIO()

	pin,err := embd.NewDigitalPin(17)

	if err != nil {
		fmt.Println(err)
	}

	pin.SetDirection(embd.In)
	pin.PullUp()

	for {

		in,err := pin.Read()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(in)
		time.Sleep(1000 * time.Millisecond)
	}

}


