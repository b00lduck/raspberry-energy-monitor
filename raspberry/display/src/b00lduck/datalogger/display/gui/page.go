package gui
import (
	"image"
	"image/draw"
)

type Page struct {
	background image.Image
	buttons []*Button
}

func NewPage() (page *Page) {
	page = new(Page)
	page.buttons = make([]*Button, 0)
	return
}

func (page *Page) Draw(target *draw.Image) {

	if page.background != nil {
		draw.Draw(*target, (*target).Bounds(), page.background, image.ZP, draw.Src)
	}

	// draw all buttons
	for b := range page.buttons {
		page.buttons[b].Draw(*target)
	}

}

func (page *Page) AddButton(img image.Image, x int, y int) *Button {
	newButton := NewButton(img, x, y)
	page.buttons = append(page.buttons, newButton)
	return newButton
}

func (page *Page) SetBackground(img image.Image) {
	page.background = img
}