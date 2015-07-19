package pages
import (
	"fmt"
	"image/draw"
	"image"
)

type GasPage struct {
	BasePage
	img image.Image
}

func CreateGasPage() Page {

	arrowUp := LoadImage("arrow_up.gif")
	arrowDown := LoadImage("arrow_down.gif")

	gasPage := *new(GasPage)
	gasPage.BasePage = NewBasePage()

	for i := 0; i < 8; i ++ {
		gasPage.BasePage.AddButton(arrowUp, 20 + i * 35, 60 , func() {
			fmt.Printf("Digit %d >UP< pressed", i)
		})
		gasPage.BasePage.AddButton(arrowDown, 20 + i * 35, 150 , func() {
			fmt.Printf("Digit %d >DOWN< pressed", i)
		})
	}

	gasPage.img = LoadImage("count_digits_grey.png")

	return gasPage

}

func (p GasPage) Draw(target *draw.Image) {

	p.BaseDraw(target)

	rect := image.Rect(35, 100, 285, 180)

	draw.Draw(*target, rect, p.img, image.ZP, draw.Over)

}
