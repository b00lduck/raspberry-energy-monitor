package framebuffer

import (
	"os"
	"syscall"
	"b00lduck/datalogger/display/tools"
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

func (f Framebuffer) convertToRgba(msb, lsb uint8) color.Color {
	val := uint16(lsb << 8) + uint16(msb)
	r := (val & 0x001F) << 11
	g := (val & 0x07E0) << 5
	b := (val & 0xF800)
	a := uint16(0xFFFF)
	return color.RGBA64{r,g,b,a}
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