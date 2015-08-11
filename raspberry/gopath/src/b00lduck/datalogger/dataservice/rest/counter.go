package rest

import (
	"github.com/gocraft/web"
	"b00lduck/datalogger/dataservice/orm"
	"time"
	"fmt"
	"github.com/jinzhu/gorm"
)

func (c *Context) CounterHandler(rw web.ResponseWriter, req *web.Request) {
	var counters []orm.Counter
	db.Find(&counters)
	marshal(rw, counters)
}

func (c *Context) CounterByIdHandler(rw web.ResponseWriter, req *web.Request) {

	id,err := parseUintPathParameter(rw, req, "id")
	if (err != nil) {
		return
	}

	var counter orm.Counter
	db.Preload("CounterEvents").Where("id = ?", id).First(&counter)
	marshal(rw, counter)
}

func (c *Context) CounterByIdTickHandler(rw web.ResponseWriter, req *web.Request) {

	id,err := parseUintPathParameter(rw, req, "id")
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

func (c *Context) CounterByIdEventsHandler(rw web.ResponseWriter, req *web.Request) {

	id,err := parseUintPathParameter(rw, req, "id")
	if (err != nil) {
		return
	}

	start,err := c.parseUintQueryParameter(rw, "start")
	if (err != nil) {
		return
	}

	end,err := c.parseUintQueryParameter(rw, "end")
	if (err != nil) {
		return
	}

	var counterEvents []orm.CounterEvent
	var q *gorm.DB

	switch {
	case start > 0 && end > 0:
		q = db.Where("counter_id = ? AND timestamp >= ? AND timestamp <= ?", id, start, end)
	case end > 0:
		q = db.Where("counter_id = ? AND timestamp <= ?", id, end)
	case start > 0:
		q = db.Where("counter_id = ? AND timestamp >= ?", id, start)
	default:
		q = db.Where("counter_id = ?", id)
	}

	q.Order("timestamp").Find(&counterEvents)

	marshal(rw, counterEvents)
}
