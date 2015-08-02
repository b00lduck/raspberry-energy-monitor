package touchscreen

import (
	"os"
	"syscall"
	"encoding/binary"
	"b00lduck/datalogger/display/tools"
)

const (
	TSEVENT_NULL = 0
	TSEVENT_PUSH = 1
	TSEVENT_RELEASE = 2
)

type InputEvent struct {
	Time syscall.Timeval
	Type uint16
	Code uint16
	Value int32
}

type TouchscreenEvent struct {
	Type uint16
	X,Y int32
}

type Touchscreen struct {
	file* os.File
	Event chan TouchscreenEvent
	tempEvent TouchscreenEvent
}

func (f *Touchscreen) Open(device string) {
	file, err := os.OpenFile(device, os.O_RDONLY, 0)
	tools.ErrorCheck(err)
	f.file = file
}

func (f *Touchscreen) Close() {
	f.file.Close()
}

func (f *Touchscreen) Run() {
	f.Event = make(chan TouchscreenEvent, 64)

	pushed := false

	for {

		inputEvent := InputEvent{}
		err := binary.Read(f.file, binary.LittleEndian, &inputEvent)
		tools.ErrorCheck(err)

		switch {
		case inputEvent.Type == 1 && inputEvent.Value == 1:
			f.tempEvent = TouchscreenEvent{TSEVENT_PUSH, f.tempEvent.X, f.tempEvent.Y}
		case inputEvent.Type == 1 && inputEvent.Value == 0:
			f.tempEvent = TouchscreenEvent{TSEVENT_RELEASE, f.tempEvent.X, f.tempEvent.Y}
		case inputEvent.Type == 3 && inputEvent.Code == 0:
			f.tempEvent.Y = 250 - int32(float32(inputEvent.Value) / 3750.0 * 240.0);
		case inputEvent.Type == 3 && inputEvent.Code == 1:
			f.tempEvent.X = int32(float32(inputEvent.Value) / 3827.0 * 320.0) - 13;
		case inputEvent.Type == 0:
			switch {
			    case f.tempEvent.Type == TSEVENT_PUSH && !pushed:
				    pushed = true
      				    f.Event <- f.tempEvent
			    case f.tempEvent.Type == TSEVENT_RELEASE:
				    pushed = false
      				    f.Event <- f.tempEvent
			}
		}

	}
}

