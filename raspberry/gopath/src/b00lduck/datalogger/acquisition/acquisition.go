package main

import (
	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/all"
	"fmt"
	"time"
)


func main() {

    if err := embd.InitGPIO(); err != nil {
	panic(err)
    }
    defer embd.CloseGPIO()

    led, err := embd.NewDigitalPin(17)
    if err != nil {
	panic(err)
    }
    defer led.Close()

    if err := led.SetDirection(embd.In); err != nil {
	panic(err)
    }

//    if err := led.PullUp(); err != nil {
//	panic(err)
//    }

    for {
	
        pin,err := led.Read()
	if  err != nil {
    	    panic(err)
	}
	fmt.Println(pin)

	time.Sleep(1 * time.Second)

    }
}


