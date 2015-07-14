package framebuffer

import (
	"os"
	"syscall"
	"b00lduck/datalogger/display/errorcheck"
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

	file, err := os.OpenFile(device, os.O_RDWR, 0)
	errorcheck.Check(err)

	data, err := syscall.Mmap(int(file.Fd()), 0, screensize,
		syscall.PROT_READ | syscall.PROT_WRITE, syscall.MAP_SHARED)
	errorcheck.Check(err)

	f.file = file
	f.data = data
}

func (f *Framebuffer) Close() {
	f.file.Close()
}

func (f *Framebuffer) Data() []byte {
	return f.data
}