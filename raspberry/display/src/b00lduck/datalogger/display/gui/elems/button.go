package elems
import (
	"image"
	"image/draw"
)

type Button struct {
	img *PosImage
	Action func()
	IsMenuButton bool
	ChangeToPage string
}

func (b *Button) Draw(target draw.Image) {
	b.img.draw(target)
}

func NewButton(img image.Image, x, y int, action func()) *Button {
	newButton := Button { NewPosImage(img, x, y), action, false, ""}
	return &newButton
}

func NewMenuButton(img image.Image, x, y int, page string) *Button {
	newButton := Button { NewPosImage(img, x, y), nil, true, page}
	return &newButton
}

func (b *Button) IsHitBy(p image.Point) bool {

	min := b.img.Bounds().Min
	max := b.img.Bounds().Max

	if (p.X > min.X) && (p.X < max.X) && (p.Y > min.Y) && (p.Y < max.Y) {
		return true
	}
	return false
}


