package rest

import (
	"github.com/gocraft/web"
	"b00lduck/datalogger/dataservice/orm"
	"github.com/jinzhu/gorm"
	"net/http"
	"fmt"
	"io/ioutil"
)

// Get all counters
func (c *Context) CounterHandler(rw web.ResponseWriter, req *web.Request) {
	var counters []orm.Counter
	db.Find(&counters)
	marshal(rw, counters)
}

// Get specific counter by id
func (c *Context) CounterByIdHandler(rw web.ResponseWriter, req *web.Request) {

	id,err := parseUintPathParameter(rw, req, "id")
	if (err != nil) {
		return
	}

	var counter orm.Counter
	db.First(&counter, id)

	if (counter.ID == 0) {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("Counter not found"))
		return
	}

	marshal(rw, counter)
}

// Tick counter by id
func (c *Context) CounterByIdTickHandler(rw web.ResponseWriter, req *web.Request) {

	id,err := parseUintPathParameter(rw, req, "id")
	if (err != nil) {
		return
	}

	var counter orm.Counter
	db.First(&counter, id)

	if (counter.ID == 0) {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("Counter not found"))
		return
	}

	counterEvent := orm.NewTickCounterEvent(counter)
	db.Create(&counterEvent)

	counter.Reading = counterEvent.Reading
	db.Save(counter)

	marshal(rw, counterEvent)
}

// Correct counter by id
func (c *Context) CounterByIdCorrectHandler(rw web.ResponseWriter, req *web.Request) {

	id,err := parseUintPathParameter(rw, req, "id")
	if (err != nil) {
		return
	}

	var counter orm.Counter
	db.First(&counter, id)

	if (counter.ID == 0) {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("Counter not found"))
		return
	}

	hah, err := ioutil.ReadAll(req.Body);
	fmt.Println(hah)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Error reading body"))
		return
	}

	newReading,err := parseUintFromString(string(hah))
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Malformed value"))
		return
	}

	delta := int64(newReading) - int64(counter.Reading)
	counter.Reading = newReading
	db.Save(counter)

	counterEvent := orm.NewAbsCorrCounterEvent(counter, delta)

	db.Create(&counterEvent)
	marshal(rw, counterEvent)
}

// Get counter events in a optionally given time range
// Query parameters: start,end
func (c *Context) CounterByIdEventsHandler(rw web.ResponseWriter, req *web.Request) {

	id,err := parseUintPathParameter(rw, req, "id")
	if (err != nil) {
		return
	}

	var counter orm.Counter
	db.First(&counter, id)

	if (counter.ID == 0) {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("Counter not found"))
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

	lastEvent := orm.NewLastCounterEvent(counter)
	if lastEvent.Timestamp > end {
		counterEvents = append(counterEvents, lastEvent)
	}

	marshal(rw, counterEvents)
}
