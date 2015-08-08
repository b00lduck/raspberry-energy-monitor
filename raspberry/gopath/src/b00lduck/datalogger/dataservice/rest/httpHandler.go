package rest
import (
	"net/http"
	"log"
	"b00lduck/datalogger/dataservice/rest/context"
)

type HttpHandler struct {
	*context.Context
	h func(*context.Context, http.ResponseWriter, *http.Request) (int, error)
}

func (handler HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	status, err := handler.h(handler.Context, w, r)

	if err != nil {
		log.Printf("HTTP %d: %q", status, err)
		switch status {
		case http.StatusNotFound:
			http.NotFound(w, r)
		case http.StatusInternalServerError:
			http.Error(w, http.StatusText(status), status)
		default:
			http.Error(w, http.StatusText(status), status)
		}
	}
}