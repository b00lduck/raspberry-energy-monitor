package gui
import (
	"image/draw"
	"b00lduck/datalogger/display/touchscreen"
	"fmt"
	"time"
	"b00lduck/datalogger/display/gui/pages"
	"image"
)

const (
	MAIN_PAGE_NAME = "MAIN"
)

type Gui struct {
	mainPage pages.Page
	pages map[string]pages.Page
	activePageName string
	target *draw.Image
	touchscreen *touchscreen.Touchscreen
	dirty bool
}

func NewGui(target draw.Image, touchscreen *touchscreen.Touchscreen) *Gui {
	newGui := new(Gui)
	newGui.pages = make (map[string]pages.Page,0)
	newGui.activePageName = ""
	newGui.target = &target
	newGui.touchscreen = touchscreen
	newGui.dirty = false
	return newGui
}

func (g *Gui) SetPage(name string, page pages.Page) {
	g.pages[name] = page
}

func (g *Gui) SetMainPage(page pages.Page) {
	g.pages[MAIN_PAGE_NAME] = page
}

func (g *Gui) SelectPage(name string) pages.Page {

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
	for i := range page.Buttons() {
		button := page.Buttons()[i]
		if button.IsHitBy(image.Pt(int(e.X), int(e.Y))) {
			if button.IsMenuButton {
				g.SelectPage(button.ChangeToPage)
			} else {
				button.Action()
			}
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

	oldEvent := touchscreen.TouchscreenEvent{touchscreen.TSEVENT_NULL, 0,0}

	for {

		select {
		case e := <- *tsEvent:
			if e.Type == touchscreen.TSEVENT_PUSH {
				if oldEvent != e {
					g.processButtonsOfPage(e, MAIN_PAGE_NAME)
					g.processButtonsOfPage(e, g.activePageName)
					oldEvent = e
				}
			}
		default:
			if (g.dirty) {
				g.drawPage(MAIN_PAGE_NAME)
				g.drawPage(g.activePageName)
				g.dirty = false
			} else {
				time.Sleep(25 * time.Millisecond)
			}
		}

	}
}