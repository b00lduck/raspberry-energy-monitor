package rest
import (
	"net/http"
	"log"
	"b00lduck/tools"
	"encoding/json"
	"b00lduck/datalogger/dataservice/orm"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Webserver struct {
	db *gorm.DB
}

func NewWebserver(db *gorm.DB) *Webserver {
	newWebserver := new(Webserver)
	newWebserver.db = db
	return newWebserver
}

func (w *Webserver) Run() {
	http.HandleFunc("/counter", w.counterHandler)
	log.Println("Serving counter at /counter")

	e := http.ListenAndServe(":8080", nil)
	tools.ErrorCheck(e)
}

func (w *Webserver) counterHandler(writer http.ResponseWriter, r *http.Request) {

	buffer, err := json.Marshal(orm.Counter{Name:"TEST", Unit:"TESTU"})
	tools.ErrorCheck(err)

	writer.Header().Set("Content-Type", "application/json")
	_, err = writer.Write(buffer)
	tools.ErrorCheck(err)
}
