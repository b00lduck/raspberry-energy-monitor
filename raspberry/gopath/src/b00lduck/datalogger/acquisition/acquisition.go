package main

import (
	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/all"
	"fmt"
	"time"
	"net/http"
	"strings"
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
			count = 0
			count()
		}

		time.Sleep(10 * time.Millisecond)

    }
}

func count() {

	client := &http.Client{}
	request, err := http.NewRequest("POST", "http://localhost:8080/counter/1/tick", strings.NewReader(""))
	if err != nil {
		fmt.Println("Error creating tick request to dataservice")
		fmt.Println(err)
		return
	}
	request.ContentLength = 0
	_, err = client.Do(request)
	if err != nil {
		fmt.Println("Error sending tick request to dataservice")
		fmt.Println(err)
	}

}

