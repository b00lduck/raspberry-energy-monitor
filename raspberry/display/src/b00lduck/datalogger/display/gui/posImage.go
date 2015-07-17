package gui
import (
	"image"
	"image/draw"
)

type PosImage struct {
	img image.Image
	x int
	y int
}

func (p *PosImage) draw (target draw.Image) {
	bounds := p.img.Bounds().Max
	rect := image.Rect(p.x, p.y, p.x + bounds.X, p.y + bounds.Y)
	draw.Draw(target, rect, p.img, image.ZP, draw.Over)
}

func NewPosImage(img image.Image, x int, y int) *PosImage {
	newPosImage := PosImage{img, x, y }
	return &newPosImage
}
