package main

import (
	"os"
	"time"
	"b00lduck/datalogger/display/framebuffer"
	"b00lduck/datalogger/display/touchscreen"
	"image"
	"image/jpeg"
	"b00lduck/datalogger/display/errorcheck"
	"strings"
	"image/png"
	"image/gif"
	"b00lduck/datalogger/display/gui"
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

	cat := loadImage("cats-q-c-320-240-3.jpg")
	arrowUp := loadImage("arrow_up.gif")
	arrowDown := loadImage("arrow_down.gif")

	gui := gui.NewGui()

	gui.SetBackground(cat)
	for i := 0; i < 8; i ++ {
		gui.AddButton(arrowUp, 20 + i * 35, 60 )
		gui.AddButton(arrowDown, 20 + i * 35, 140 )
	}

	go gui.Run(displayBuffer, &ts.Event)

	for {
		drawDisplay(data)
		time.Sleep(100 * time.Millisecond)
	}

}

func loadImage(filename string) image.Image {

	f, err := os.Open("images/" + filename)
	errorcheck.Check(err)

	var img image.Image = nil

	lowerFilename := strings.ToLower(filename)
	switch {
		case strings.HasSuffix(lowerFilename, ".jpg"):
			img, err = jpeg.Decode(f)
		case strings.HasSuffix(lowerFilename, ".png"):
			img, err = png.Decode(f)
		case strings.HasSuffix(lowerFilename, ".gif"):
			img, err = gif.Decode(f)
	}

	errorcheck.Check(err)

	return img
}
