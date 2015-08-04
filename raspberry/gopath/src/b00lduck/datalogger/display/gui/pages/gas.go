package pages
import (
	"fmt"
	"image/draw"
	"image"
	"strconv"
)

type GasPage struct {
	BasePage
	imgBlk image.Image
	imgRed image.Image
	Counter int32
}

func CreateGasPage() Page {

	fmt.Println("CREATE GAS PAGE")

	arrowUp := LoadImage("arrow_up.gif")
	arrowDown := LoadImage("arrow_down.gif")

	gasPage := *new(GasPage)
	gasPage.BasePage = NewBasePage()

	gasPage.Counter = 22374090

	for i := 0; i < 8; i ++ {
		gasPage.BasePage.AddButton(arrowUp, 20 + i * 35, 60 , gasPage.incHandler(i))
		gasPage.BasePage.AddButton(arrowDown, 20 + i * 35, 150 , gasPage.decHandler(i))
	}

	gasPage.imgBlk = LoadImage("count_digits_grey.png")
	gasPage.imgRed = LoadImage("count_digits_red.png")

	return &gasPage

}

func (p *GasPage) incHandler(i int) func() {
	return func() {
		fmt.Printf("Digit %d >UP< pressed\n", i)
		p.changeCounter(i, true)
	}
}

func (p *GasPage) decHandler(i int) func() {
	return func() {
		fmt.Printf("Digit %d >DOWN< pressed\n", i)
		p.changeCounter(i, false)
	}
}

func (p GasPage) DrawDigit(target draw.Image, src image.Image, digit uint8, pos image.Point) {

	digitWidth := 25
	digitHeight := 40

	targetRect := image.Rect(pos.X, pos.Y, pos.X + digitWidth, pos.Y + digitHeight)

	sourcePos := image.Pt(digitWidth * int(digit), 0)

	draw.Draw(target, targetRect, src, sourcePos, draw.Over)

}

func (p *GasPage) Draw(target *draw.Image) {

	p.BaseDraw(target)

	cstr := fmt.Sprintf("%08d", p.Counter)

	for i := 0; i < 8; i++ {
		pint, _ := strconv.ParseUint(string(cstr[i]), 10, 8)
		if i < 5 {
			p.DrawDigit(*target, p.imgBlk, uint8(pint), image.Pt(20 + 36 * i, 100))
		} else {
			p.DrawDigit(*target, p.imgRed, uint8(pint), image.Pt(20 + 36 * i, 100))
		}

	}

}

func (p *GasPage) Process() bool {
	p.Counter += 10
	return true
}

func (p *GasPage) changeCounter(digit int, direction bool) {

	cstr := fmt.Sprintf("%08d", p.Counter)

	pint, _ := strconv.ParseUint(string(cstr[digit]), 10, 8)

	factor := pow10(7 - digit)

	if direction {
		if pint == 9 {
			p.Counter -= factor * 9
		} else {
			p.Counter += factor
		}
	} else {
		if pint == 0 {
			p.Counter += factor * 9
		} else {
			p.Counter -= factor
		}
	}

	*(p.BasePage.DirtyChan) <- true

}

func pow10(n int) (ret int32) {
	ret = 1
	for ;n > 0;n-- {
		ret *= 10
	}
	return
}
