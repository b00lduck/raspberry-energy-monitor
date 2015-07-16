package main

import (
	"os"
	"time"
	"image/draw"
	"b00lduck/datalogger/display/framebuffer"
	"b00lduck/datalogger/display/touchscreen"
	"image"
	"image/jpeg"
	"b00lduck/datalogger/display/errorcheck"
)

var mode int = 3

var displayBuffer* image.RGBA

func drawDisplay(data []byte) {

	srcCount := 0
	targetCount := 0

	for srcCount < 240*320*4 {
		r := displayBuffer.Pix[srcCount]
		srcCount++;
		g := displayBuffer.Pix[srcCount]
		srcCount++;
		b := displayBuffer.Pix[srcCount]
		srcCount++;
		srcCount++;

		r8 := uint16(r >> 3)
		g8 := uint16(g >> 2)
		b8 := uint16(b >> 3)

		out := r8 << 11 + g8 << 5 + b8

		outl := uint8(out >> 8)
		outh := uint8(out & 0xff)

		data[targetCount] = outh
		targetCount++
		data[targetCount] = outl
		targetCount++
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

	displayBuffer = image.NewRGBA(image.Rect(0, 0, 320, 240))

	f, err := os.Open("images/cats-q-c-320-240-3.jpg")
	errorcheck.Check(err)

	cat, err := jpeg.Decode(f)
	errorcheck.Check(err)

	draw.Draw(displayBuffer, displayBuffer.Bounds(), cat, image.ZP, draw.Src)



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
				drawDisplay(data)
				time.Sleep(20 * time.Millisecond)
		}
	}

}
