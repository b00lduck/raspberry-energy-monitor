package pages
import (
	"image"
	"image/draw"
	"os"
	"b00lduck/tools"
	"strings"
	"image/jpeg"
	"image/png"
	"image/gif"
	"b00lduck/datalogger/display/gui/elems"
)

type Page interface {
	Draw(*draw.Image)
	Process() bool
	Buttons() []*elems.Button
	SetDirtyChan(*chan bool)
}

type BasePage struct {
	background image.Image
	buttons []*elems.Button
	DirtyChan *chan bool
}

func NewBasePage() (page BasePage) {
	page = *new(BasePage)
	page.buttons = make([]*elems.Button, 0)
	return
}

func (page *BasePage) SetDirtyChan(dirtyChan *chan bool) {
	page.DirtyChan = dirtyChan
}

func (page BasePage) Draw(target *draw.Image) {
	page.BaseDraw(target)
}

func (page BasePage) BaseDraw(target *draw.Image) {

	if page.background != nil {
		draw.Draw(*target, (*target).Bounds(), page.background, image.ZP, draw.Src)
	}

	// draw all buttons
	for b := range page.buttons {
		page.buttons[b].Draw(*target)
	}

}

func (page BasePage) Process() bool {
	return page.BaseProcess()
}

func (page BasePage) BaseProcess() bool{
	return false
}

func (page *BasePage) AddButton(img image.Image, x, y int, action func()) *elems.Button {
	newButton := elems.NewButton(img, x, y, action)
	page.buttons = append(page.buttons, newButton)
	return newButton
}

func (page *BasePage) AddMenuButton(img image.Image, x, y int, newPage string) *elems.Button {
	newButton := elems.NewMenuButton(img, x, y, newPage)
	page.buttons = append(page.buttons, newButton)
	return newButton
}

func (page *BasePage) SetBackground(img image.Image) {
	page.background = img
}

func (page BasePage) Buttons() []*elems.Button {
	return page.buttons
}

func LoadImage(filename string) image.Image {

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