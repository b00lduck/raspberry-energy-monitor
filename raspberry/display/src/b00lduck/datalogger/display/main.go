package main

import (
	"os"
	"time"
	"b00lduck/datalogger/display/framebuffer"
	"b00lduck/datalogger/display/touchscreen"
	"image"
	"image/jpeg"
	"strings"
	"image/png"
	"image/gif"
	"b00lduck/datalogger/display/gui"
	"b00lduck/datalogger/display/tools"
	"b00lduck/datalogger/display/i2c"
	"fmt"
)

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

	mm, err := i2c.New()
	tools.ErrorCheck(err)

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

	//go gui.Run(displayBuffer, &ts.Event)

	gui.Draw(displayBuffer)
	drawDisplay(data)

	for {
		//drawDisplay(data)
		time.Sleep(100 * time.Millisecond)

		x := uint16(mm.Read(8)) << 8 + mm.Read(9)
		y := uint16(mm.Read(10)) << 8 + mm.Read(11)
		z := uint16(mm.Read(12)) << 8 + mm.Read(13)

		fmt.Printf("X:%d Y:%d Z:%d", x, y, z)
	}

}

func loadImage(filename string) image.Image {

	f, err := os.Open("images/" + filename)
	tools.ErrorCheck(err)

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

	tools.ErrorCheck(err)

	return img
}
