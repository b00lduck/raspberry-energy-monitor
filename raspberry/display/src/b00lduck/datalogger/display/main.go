package main

import (
	"os"
	"time"
	"math/rand"
	"b00lduck/datalogger/display/framebuffer"
	"b00lduck/datalogger/display/touchscreen"
)

var mode int

func draw(data []byte) {

	switch mode {
	case 0:
		for i := range data {
			data[i] = byte(rand.Int())
		}
	case 1:
		for i := range data {
			data[i] = 0
		}
	case 2:
		for i := range data {
			data[i] = 255
		}
	case 3:
		for i := range data {
			data[i] = byte(i)
		}
	}

}

func main() {

	fb := framebuffer.Framebuffer{}
	fb.Open(os.Args[1])
	defer fb.Close()
	data := fb.Data()

	ts := touchscreen.Touchscreen{}
	ts.Open(os.Args[2])
	defer ts.Close()
	go ts.Run()

	for {
		select {
			case event := <-ts.Event:
				if event.Type == touchscreen.TSEVENT_PUSH {
					if event.X < 160 {
						mode = 0
					} else {
						mode = 1
					}
					if event.Y > 120 {
						mode += 2;
					}
				}
			default:
				draw(data)
				time.Sleep(20 * time.Millisecond)
		}
	}

}
