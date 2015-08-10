package rest
import (
	"github.com/gocraft/web"
	"b00lduck/datalogger/dataservice/orm"
	"encoding/json"
	"b00lduck/tools"
	"strconv"
	"log"
)

func (c *Context) CounterHandler(rw web.ResponseWriter, req *web.Request) {

	var counters []orm.Counter

	db.Find(&counters)

	buffer, err := json.Marshal(counters)
	tools.ErrorCheck(err)

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err = rw.Write(buffer)
	tools.ErrorCheck(err)

}

func (c *Context) CounterHandlerId(rw web.ResponseWriter, req *web.Request) {

	id, _ := strconv.ParseUint(req.PathParams["id"], 10, 8)

	log.Print("Counter ID:", req.PathParams["id"])

	var counter orm.Counter
	db.Where("id = ?", id).First(&counter)

	buffer, err := json.Marshal(counter)
	tools.ErrorCheck(err)

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err = rw.Write(buffer)
	tools.ErrorCheck(err)

}