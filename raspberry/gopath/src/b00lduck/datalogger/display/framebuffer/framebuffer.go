package framebuffer

import (
	"os"
	"syscall"
	"b00lduck/tools"
	"image/color"
	"image"
	"fmt"
)

const resx = 320
const resy = 240
const depth = 16
const screensize = resx * resy * depth / 8

type Framebuffer struct {
	file* os.File
	data []byte
}

func (f *Framebuffer) Open(device string) {

	fmt.Printf("initializing framebuffer on device %s\n", device)

	file, err := os.OpenFile(device, os.O_RDWR, 0)
	tools.ErrorCheck(err)

	data, err := syscall.Mmap(int(file.Fd()), 0, screensize,
		syscall.PROT_READ | syscall.PROT_WRITE, syscall.MAP_SHARED)
	tools.ErrorCheck(err)

	f.file = file
	f.data = data
}

func (f *Framebuffer) Close() {
	f.file.Close()
}

func (f *Framebuffer) Data() []byte {
	return f.data
}

func (f Framebuffer) convertToFb(c color.Color) (msb,lsb uint8) {
	r,g,b,_ := c.RGBA()
	out := uint16(r >> 11) << 11 + uint16(g >> 10) << 5 + uint16(b >> 11)
	lsb = uint8(out >> 8)
	msb = uint8(out & 0xff)
	return
}

func (f Framebuffer) Set(x, y int, c color.Color) {
	offset := x * 2 + y * 640
	f.data[offset], f.data[offset + 1] = f.convertToFb(c)
}

func (f Framebuffer) convertToRgba(lsb, msb uint8) color.Color {
	val := uint16(msb) << 8 + uint16(lsb)
	b := uint8 ((val & 0x001F) << 3)
	g := uint8 ((val & 0x07E0) >> 3)
	r := uint8 ((val & 0xF800) >> 8)
	a := uint8 (0xFF)
	return color.RGBA{r,g,b,a}
}

func (f Framebuffer) At(x, y int) color.Color {
	offset := x * 2 + y * 640
	return f.convertToRgba(f.data[offset], f.data[offset + 1])
}

func (f Framebuffer) Bounds() image.Rectangle {
	return image.Rectangle{image.ZP, image.Point{320,240}}
}

func (f Framebuffer) ColorModel() color.Model {
	return color.RGBAModel
}