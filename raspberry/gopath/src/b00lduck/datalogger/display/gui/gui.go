package gui
import (
	"image/draw"
	"b00lduck/datalogger/display/touchscreen"
	"fmt"
	"b00lduck/datalogger/display/gui/pages"
	"image"
)

type Gui struct {
	mainPage *pages.Page
	clearPage pages.Page
	pages map[string]*pages.Page
	activePageName string
	target *draw.Image
	touchscreen *touchscreen.Touchscreen
	dirty chan bool
	Bounds image.Rectangle
	timeout int32
}

// Screensaver
const DISPLAY_TIMEOUT = 5

func NewGui(target draw.Image, touchscreen *touchscreen.Touchscreen) *Gui {
	newGui := new(Gui)
	newGui.pages = make (map[string]*pages.Page,0)
	newGui.clearPage = pages.NewClearPage()
	newGui.activePageName = ""
	newGui.target = &target
	newGui.touchscreen = touchscreen
	newGui.dirty = make(chan bool, 64)
	newGui.Bounds = image.Rect(0,0,320,240)
	return newGui
}

func (g *Gui) SetPage(name string, page pages.Page) {
	g.pages[name] = &page
	page.SetDirtyChan(&g.dirty)
}

func (g *Gui) SetMainPage(page pages.Page) {
	g.mainPage = &page
	page.SetDirtyChan(&g.dirty)
}

func (g *Gui) SelectPage(name string) {

	if g.activePageName == name {
		fmt.Println("page stays page " + name)
		return
	}

	g.dirty <- true

	if g.pages[name] != nil {
		fmt.Println("success selecting page " + name)
		g.activePageName = name
		return
	}
	fmt.Println("fail selecting page " + name)
	g.activePageName = ""
}

func (g *Gui) processButtonsOfPage(e touchscreen.TouchscreenEvent, page *pages.Page) {
	if page == nil {
		return
	}
	page2 := *page
	for i := range page2.Buttons() {
		button := page2.Buttons()[i]
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
	page := *g.pages[name]
	page.Draw(g.target)
}

func (g *Gui) Process() {

	if (g.timeout > DISPLAY_TIMEOUT) {
		return
	}

	g.timeout += 1

	if (g.timeout == DISPLAY_TIMEOUT + 1) {
		g.dirty <- true
		return
	}

	mp := *g.mainPage
	dirty := mp.Process()

	for i := range g.pages {
		p := *g.pages[i]
		if p.Process() {
			dirty = true
		}

	}

	if dirty {
		g.dirty <- true
	}
}

func (g *Gui) Run(tsEvent *chan touchscreen.TouchscreenEvent) {

	oldEvent := touchscreen.TouchscreenEvent{touchscreen.TSEVENT_NULL, 0,0}

	for {
		select {
		case e := <- *tsEvent:
			if e.Type == touchscreen.TSEVENT_PUSH {
				if oldEvent != e {
					if (g.timeout > DISPLAY_TIMEOUT) {
						g.dirty <- true
					} else {
						g.processButtonsOfPage(e, g.mainPage)
						g.processButtonsOfPage(e, g.pages[g.activePageName])
					}
					g.timeout = 0
					oldEvent = e
				}
			} else {
				oldEvent = touchscreen.TouchscreenEvent{touchscreen.TSEVENT_NULL, 0,0}
			}

		case <- g.dirty:

			if g.timeout > DISPLAY_TIMEOUT {
				draw.Draw(*g.target, g.Bounds, image.NewUniform(image.Black), image.ZP, draw.Src)
			} else {
				doubleBuffer := draw.Image(image.NewRGBA(g.Bounds))

				mainPage := g.mainPage
				if (mainPage != nil) {
					(*mainPage).Draw(&doubleBuffer)
				}

				currentPage := g.pages[g.activePageName]
				if (currentPage != nil) {
					(*currentPage).Draw(&doubleBuffer)
				}

				draw.Draw(*g.target, doubleBuffer.Bounds(), doubleBuffer, image.ZP, draw.Src)
			}
		}
	}
}