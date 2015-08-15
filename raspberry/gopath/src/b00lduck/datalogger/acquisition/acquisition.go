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

	embd.SetDirection(17, embd.In)
	embd.PullUp(17)

	for {

		in,_ := embd.DigitalRead(17)

		fmt.Println(in)

		time.Sleep(1000 * time.Millisecond)

	}

}


