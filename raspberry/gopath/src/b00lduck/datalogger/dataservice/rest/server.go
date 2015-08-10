package rest
import (
	"github.com/jinzhu/gorm"
	"net/http"
	"b00lduck/tools"
	"github.com/gocraft/web"
	"encoding/json"
	"log"
	"strconv"
)

var db *gorm.DB

type Context struct {
}

func StartServer(database *gorm.DB) {

	db = database

	router := web.New(Context{}).
	Middleware(web.LoggerMiddleware).

	Get ("/counter", 			(*Context).CounterHandler).
	Get ("/counter/:id", 		(*Context).CounterByIdHandler).
	Post("/counter/:id/tick", 	(*Context).CounterByIdTickHandler).
	Get ("/counterEvent", 		(*Context).CounterEventHandler)

	e := http.ListenAndServe(":8080", router)
	tools.ErrorCheck(e)
}

func marshal(rw web.ResponseWriter, v interface{}) {

	buffer, err := json.Marshal(v)
	if err != nil {
		log.Println("Error marshalling JSON")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")

	if _, err = rw.Write(buffer); err != nil {
		log.Println("Error writing JSON to buffer")
		rw.WriteHeader(http.StatusInternalServerError)
	}

}

func parseUintParameter(rw web.ResponseWriter, req *web.Request, name string) (id uint64, err error) {

	sId := req.PathParams[name]
	id, err = strconv.ParseUint(sId, 10, 8)

	if (err != nil) {
		log.Println("Error parsing id: " + sId)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Malformed parameter 'id'"))
	}

	return
}
