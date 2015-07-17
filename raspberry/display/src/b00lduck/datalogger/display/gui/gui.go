package gui
import (
	"image"
	"image/color"
	"image/draw"
	"time"
	"b00lduck/datalogger/display/touchscreen"
	"fmt"
)

type Gui struct {
	background image.Image
	buttons []*Button
}

func (g *Gui) Draw(target draw.Image) {

	if g.background == nil {
		// no background
		black := image.Uniform{color.Black}
		draw.Draw(target, target.Bounds(), &black, image.ZP, draw.Src)
	} else {
		// draw background
		draw.Draw(target, target.Bounds(), g.background, image.ZP, draw.Src)
	}

	// draw all buttons
	for b := range g.buttons {
		g.buttons[b].Draw(target)
	}
}

func (g *Gui) AddButton(img image.Image, x int, y int) {
	newButton := NewButton(img, x, y)
	g.buttons = append(g.buttons, newButton)
}

func (g *Gui) SetBackground(img image.Image) {
	g.background = img
}

func NewGui() *Gui {
	newGui := Gui{}
	newGui.buttons = make([]*Button, 0)
	return &newGui
}

func (g *Gui) Run(target draw.Image, event *chan touchscreen.TouchscreenEvent) {

	for {
		select {
		case e := <-*event:
			if e.Type == touchscreen.TSEVENT_PUSH {
				for i := range g.buttons {

					x := int(e.X)
					y := int(e.Y)
					min := g.buttons[i].img.Bounds().Min
					max := g.buttons[i].img.Bounds().Max

					if (x > min.X) && (x < max.X) && (y > min.Y) && (y < max.Y) {
						fmt.Printf("BUTTON %d PRESSED\n", i)
					}

				}

			}
		default:
			g.Draw(target)
			time.Sleep(100 * time.Millisecond)
		}
	}

}