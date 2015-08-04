package webserver
import (
	"net/http"
	"log"
	"image"
	"bytes"
	"strconv"
	"image/png"
	"b00lduck/datalogger/display/tools"
	"b00lduck/datalogger/display/touchscreen"
)

type Webserver struct {
	img *image.Image
	tsEvent *chan touchscreen.TouchscreenEvent
}

func NewWebserver(img image.Image, tsEvent *chan touchscreen.TouchscreenEvent) *Webserver {
	newWebserver := new(Webserver)
	newWebserver.img = &img
	newWebserver.tsEvent = tsEvent
	return newWebserver
}

func (w *Webserver) Run() {
	http.HandleFunc("/display", w.displayHandler)
	log.Println("Serving current display image at /display")

	http.HandleFunc("/click", w.clickHandler)
	log.Println("Waiting for click events at /click")

	e := http.ListenAndServe(":8081", nil)
	tools.ErrorCheck(e)
}

func (w *Webserver) displayHandler(writer http.ResponseWriter, r *http.Request) {
	buffer := new(bytes.Buffer)
	e := png.Encode(buffer, *w.img)
	tools.ErrorCheck(e)

	writer.Header().Set("Content-Type", "image/png")
	writer.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	_, e = writer.Write(buffer.Bytes())
	tools.ErrorCheck(e)
}

func (w *Webserver) clickHandler(writer http.ResponseWriter, r *http.Request) {

	x,e := readParameter(r, "x")
	if e != nil {
		http.Error(writer, "Malformed parameter y", http.StatusBadRequest)
		return
	}

	y,e := readParameter(r, "y")
	if e != nil {
		http.Error(writer, "Malformed parameter y", http.StatusBadRequest)
		return
	}

	*w.tsEvent <- touchscreen.TouchscreenEvent{touchscreen.TSEVENT_PUSH, x, y}
	*w.tsEvent <- touchscreen.TouchscreenEvent{touchscreen.TSEVENT_RELEASE, x, y}
}

func readParameter(r *http.Request, name string) (int32, error) {

	xs := r.URL.Query().Get(name)
	x,e := strconv.ParseUint(xs, 10, 16)
	if e != nil {
		return 0, e
	}
	return int32(x), nil

}