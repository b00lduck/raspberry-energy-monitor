package gui
import (
	"image/draw"
	"b00lduck/datalogger/display/touchscreen"
	"fmt"
	"time"
)

const (
	DEFAULT_PAGE_NAME = "DEFAULT"
)

type Gui struct {
	pages map[string]*Page
	activePageName string
	target *draw.Image
	touchscreen *touchscreen.Touchscreen
	butEvent chan Button
	dirty bool
}

func NewGui(target draw.Image, touchscreen *touchscreen.Touchscreen, butEvent chan Button) *Gui {
	newGui := new(Gui)
	newGui.pages = make (map[string]*Page,0)
	newGui.activePageName = ""
	newGui.target = &target
	newGui.touchscreen = touchscreen
	newGui.AddPage(DEFAULT_PAGE_NAME)
	newGui.butEvent = butEvent
	newGui.dirty = false
	return newGui
}

func (g *Gui) AddPage(name string) *Page {
	newPage := NewPage()
	g.pages[name] = newPage
	return newPage
}

func (g *Gui) GetDefaultPage() *Page {
	return g.pages[DEFAULT_PAGE_NAME]
}

func (g *Gui) SelectPage(name string) *Page {

	if g.activePageName == name {
		return g.pages[name]
	}

	g.dirty = true

	if g.pages[name] != nil {
		fmt.Println("success selecting page " + name)
		g.activePageName = name
		return g.pages[name]
	}
	fmt.Println("fail selecting page " + name)
	g.activePageName = ""
	return nil
}

func (g *Gui) processButtonsOfPage(e touchscreen.TouchscreenEvent, name string) {

	if name == "" {
		return
	}

	page := g.pages[name]

	for i := range page.buttons {

		x := int(e.X)
		y := int(e.Y)
		min := page.buttons[i].img.Bounds().Min
		max := page.buttons[i].img.Bounds().Max

		if (x > min.X) && (x < max.X) && (y > min.Y) && (y < max.Y) {
			g.butEvent <- *page.buttons[i]
		}

	}

}

func (g * Gui) drawPage(name string) {
	if name == "" {
		return
	}
	g.pages[name].Draw(g.target)
}

func (g *Gui) Run(tsEvent *chan touchscreen.TouchscreenEvent) {

	for {

		select {
		case e := <- *tsEvent:
			if e.Type == touchscreen.TSEVENT_RELEASE {
				g.processButtonsOfPage(e, DEFAULT_PAGE_NAME)
				g.processButtonsOfPage(e, g.activePageName)
			}
		default:
			if (g.dirty) {
				g.drawPage(DEFAULT_PAGE_NAME)
				g.drawPage(g.activePageName)
				g.dirty = false
			} else {
				time.Sleep(25 * time.Millisecond)
			}
		}

	}
}