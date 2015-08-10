package rest

import (
	"github.com/gocraft/web"
	"b00lduck/datalogger/dataservice/orm"
	"time"
	"fmt"
)

func (c *Context) CounterHandler(rw web.ResponseWriter, req *web.Request) {
	var counters []orm.Counter
	db.Find(&counters)
	marshal(rw, counters)
}

func (c *Context) CounterByIdHandler(rw web.ResponseWriter, req *web.Request) {

	id,err := parseUintParameter(rw, req, "id")
	if (err != nil) {
		return
	}

	var counter orm.Counter
	db.Preload("CounterEvents").Where("id = ?", id).First(&counter)
	marshal(rw, counter)
}

func (c *Context) CounterByIdTickHandler(rw web.ResponseWriter, req *web.Request) {

	id,err := parseUintParameter(rw, req, "id")
	if (err != nil) {
		return
	}

	fmt.Println(time.StampMilli)

	counterEvent := orm.CounterEvent{
		CounterID: uint(id),
		Timestamp: time.Now().UnixNano(),
		EventType: 1,
		Delta:     10,
		Reading:   12345}

	db.Create(&counterEvent)
	marshal(rw, counterEvent)
}
