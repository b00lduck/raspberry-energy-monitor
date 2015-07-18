package tools
import (
	"syscall"
)

func Ioctl(fd, cmd, arg uintptr) (err error) {
	_, _, err = syscall.Syscall6(syscall.SYS_IOCTL, fd, cmd, arg, 0, 0, 0)
	ErrorCheck(err)
	return
}

