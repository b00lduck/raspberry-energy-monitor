package webserver
import (
	"net/http"
	"log"
	"image"
	"bytes"
	"strconv"
	"image/png"
)

type Webserver struct {
	img *image.Image
}

func NewWebserver(img image.Image) *Webserver {
	newWebserver := new(Webserver)
	newWebserver.img = &img
	return newWebserver
}

func (w *Webserver) Run() {
	http.HandleFunc("/", w.displayHandler)
	log.Println("Serving current display image on localhost:8081/")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func (w *Webserver) displayHandler(writer http.ResponseWriter, r *http.Request) {
	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, *w.img); err != nil {
		log.Println("unable to encode image.")
	}

	writer.Header().Set("Content-Type", "image/png")
	writer.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := writer.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}
