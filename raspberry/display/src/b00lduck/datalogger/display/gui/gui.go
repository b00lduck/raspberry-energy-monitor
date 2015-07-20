package gui
import (
	"image/draw"
	"b00lduck/datalogger/display/touchscreen"
	"fmt"
	"time"
	"b00lduck/datalogger/display/gui/pages"
	"image"
)

type Gui struct {
	mainPage *pages.Page
	pages map[string]*pages.Page
	activePageName string
	target *draw.Image
	touchscreen *touchscreen.Touchscreen
	dirty bool
}

func NewGui(target draw.Image, touchscreen *touchscreen.Touchscreen) *Gui {
	newGui := new(Gui)
	newGui.pages = make (map[string]*pages.Page,0)
	newGui.activePageName = ""
	newGui.target = &target
	newGui.touchscreen = touchscreen
	newGui.dirty = false
	return newGui
}

func (g *Gui) SetPage(name string, page pages.Page) {
	g.pages[name] = &page
}

func (g *Gui) SetMainPage(page pages.Page) {
	g.mainPage = &page
}

func (g *Gui) SelectPage(name string) {

	if g.activePageName == name {
		fmt.Println("page stays page " + name)
		return
	}

	g.dirty = true

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

	mp := *g.mainPage
	dirty := mp.Process()

	for i := range g.pages {
		p := *g.pages[i]
		if p.Process() {
			dirty = true
		}

	}

	if dirty {
		g.dirty = true
	}

}

func (g *Gui) Run(tsEvent *chan touchscreen.TouchscreenEvent) {

	oldEvent := touchscreen.TouchscreenEvent{touchscreen.TSEVENT_NULL, 0,0}

	for {

		select {
		case e := <- *tsEvent:
			if e.Type == touchscreen.TSEVENT_PUSH {
				if oldEvent != e {
					g.processButtonsOfPage(e, g.mainPage)
					g.processButtonsOfPage(e, g.pages[g.activePageName])
					oldEvent = e
				}
			}
		default:
			if (g.dirty) {

				mainPage := g.mainPage
				if (mainPage != nil) {
					(*mainPage).Draw(g.target)
				}

				currentPage := g.pages[g.activePageName]
				if (currentPage != nil) {
					(*currentPage).Draw(g.target)
				}

				g.dirty = false
			} else {
				time.Sleep(25 * time.Millisecond)
			}
		}

	}
}