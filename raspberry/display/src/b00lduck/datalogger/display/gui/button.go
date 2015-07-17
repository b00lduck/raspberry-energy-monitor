package gui
import (
	"image"
	"image/draw"
)

type Button struct {
	img *PosImage
}

func (b *Button) Draw(target draw.Image) {
	b.img.draw(target)
}

func NewButton(img image.Image, x int, y int) *Button {
	newButton := Button { NewPosImage(img, x, y) }
	return &newButton
}