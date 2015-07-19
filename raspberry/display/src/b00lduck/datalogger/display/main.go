package main

import (
	"os"
	"b00lduck/datalogger/display/framebuffer"
	"b00lduck/datalogger/display/touchscreen"
	"b00lduck/datalogger/display/gui"
	"fmt"
	"b00lduck/datalogger/display/gui/pages"
	"time"
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

	g := gui.NewGui(fb, ts)

	g.SetMainPage(pages.CreateMainPage())
	g.SetPage("GAS_1", pages.CreateGasPage())

	g.SelectPage("GAS_1")

	go g.Run(&ts.Event)

	for {
		time.Sleep(1 * time.Second)
	}

}

