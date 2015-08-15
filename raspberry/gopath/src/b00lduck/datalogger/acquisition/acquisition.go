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

	state := false
	count := 0

    for {
        pin,err := led.Read()
		if  err != nil {
			panic(err)
		}

		if pin == 1 {
			if state == false {
				count++
			}
		} else {
			count = 0
			state = false
		}

		if count >= 3 {
			state = true
			fmt.Println("COUNT")
		}

		time.Sleep(10 * time.Millisecond)

    }
}


