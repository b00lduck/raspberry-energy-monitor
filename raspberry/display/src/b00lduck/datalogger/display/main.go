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
	"fmt"
	"b00lduck/datalogger/display/i2c/hm5883l"
//	"image/color"
	"image/color"
)

func main() {

	var mm hmc5883l.HMC5883L

	mm, err := hmc5883l.CreateHMC5883LI2c(1)
	if (err != nil) {
		mm = hmc5883l.CreateHMC5883LMock()
		fmt.Println("WARNING: HMC5883L magnetometer init failed, using mock instead. No real data will be available!")

	}

	fb := framebuffer.Framebuffer{}
	fb.Open(os.Args[1])
	defer fb.Close()
	//data := fb.Data()

	ts := touchscreen.Touchscreen{}
	ts.Open(os.Args[2])
	defer ts.Close()
	go ts.Run()

	background := loadImage("bg.png")
	//arrowUp := loadImage("arrow_up.gif")
	//arrowDown := loadImage("arrow_down.gif")

	gui := gui.NewGui()

	gui.SetBackground(background)
	/*
	for i := 0; i < 8; i ++ {
		gui.AddButton(arrowUp, 20 + i * 35, 60 )
		gui.AddButton(arrowDown, 20 + i * 35, 140 )
	}
	*/

	go gui.Run(fb, &ts.Event)

	xcount := 0

	for {
		time.Sleep(50 * time.Millisecond)
		vector, err := mm.ReadVector()
		tools.ErrorCheck(err)

		xcount++;
		if xcount >= 320 {
			xcount = 0;
		}

		tx := int(190 - (float32(vector.X) / 65535) * 140)
		ty := int(190 - (float32(vector.Y) / 65535) * 140)
		tz := int(190 - (float32(vector.Z) / 65535) * 140)

		fb.Set(xcount, tx, color.RGBA{255,0,0,255})
		fb.Set(xcount, ty, color.RGBA{0,255,0,255})
		fb.Set(xcount, tz, color.RGBA{0,0,255,255})

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
