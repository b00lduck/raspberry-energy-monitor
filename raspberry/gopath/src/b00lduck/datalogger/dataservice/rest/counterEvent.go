package rest

import (
	"github.com/gocraft/web"
	"b00lduck/datalogger/dataservice/orm"
)

func (c *Context) CounterEventHandler(rw web.ResponseWriter, req *web.Request) {
	var counterEvents []orm.CounterEvent
	db.Find(&counterEvents)
	marshal(rw, counterEvents)
}

