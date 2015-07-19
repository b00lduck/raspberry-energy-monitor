package main

import (
	"os"
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
)

func main() {

	fmt.Println("START")

	fb := new(framebuffer.Framebuffer)
	fb.Open(os.Args[1])
	defer fb.Close()

	ts := new(touchscreen.Touchscreen)
	ts.Open(os.Args[2])
	defer ts.Close()
	go ts.Run()

	background := loadImage("bg.png")
	arrowUp := loadImage("arrow_up.gif")
	arrowDown := loadImage("arrow_down.gif")

	butChan := make(chan gui.Button)

	g := gui.NewGui(fb, ts, butChan)

	defaultPage := g.GetDefaultPage()
	defaultPage.SetBackground(background)

	buttonGas := defaultPage.AddButton(arrowUp, 0, 199)
	buttonEle := defaultPage.AddButton(arrowUp, 80, 199)
	buttonWat := defaultPage.AddButton(arrowUp, 160, 199)
	buttonSys := defaultPage.AddButton(arrowUp, 240, 199)

	gas1Page := g.AddPage("GAS_1")
	for i := 0; i < 8; i ++ {
		gas1Page.AddButton(arrowUp, 20 + i * 35, 60 )
		gas1Page.AddButton(arrowDown, 20 + i * 35, 140 )
	}

	g.AddPage("ELE_1")
	g.AddPage("WAT_1")
	g.AddPage("SYS_1")


	go g.Run(&ts.Event)

	for {
		b := <- butChan

		switch b{
		case *buttonGas:
			fmt.Println("GAS")
			g.SelectPage("GAS_1")
		case *buttonEle:
			fmt.Println("ELE")
			g.SelectPage("ELE_1")
		case *buttonWat:
			fmt.Println("WAT")
			g.SelectPage("WAT_1")
		case *buttonSys:
			fmt.Println("SYS")
			g.SelectPage("SYS_1")
		}

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
