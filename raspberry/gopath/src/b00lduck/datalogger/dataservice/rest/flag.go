package rest

import (
	"github.com/gocraft/web"
	"b00lduck/datalogger/dataservice/orm"
	"net/http"
	"io/ioutil"
	"strconv"
)

// Get all flags
func (c *Context) FlagHandler(rw web.ResponseWriter, req *web.Request) {
	var flags []orm.Counter
	db.Find(&flags)
	marshal(rw, flags)
}

// Get specific flag by code
func (c *Context) FlagByCodeHandler(rw web.ResponseWriter, req *web.Request) {

	code := parseStringPathParameter(req, "code")
	var flag orm.Flag
	db.Where(&orm.Flag{Code: code}).First(&flag)

	if (flag.ID == 0) {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("Flag not found"))
		return
	}

	marshal(rw, flag)
}

// Chenge flag state by code
func (c *Context) FlagByCodeChangeStateHandler(rw web.ResponseWriter, req *web.Request) {

	code := parseStringPathParameter(req, "code")
	var flag orm.Flag
	db.Where(&orm.Flag{Code: code}).First(&flag)

	if (flag.ID == 0) {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("Flag not found"))
		return
	}

	body, err := ioutil.ReadAll(req.Body);
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Could not read body"))
		return
	}

	state, err := strconv.Atoi(string(body))
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Could not parse state"))
		return
	}

	flagState := orm.NewFlagState(flag, uint8(state))
	db.Create(&flagState)

	flag.State = uint8(state)
	flag.LastChange = flagState.Timestamp
	db.Save(flag)

	marshal(rw, flagState)
}

// Get flag states in a optionally given time range
// Query parameters: start,end
func (c *Context) FlagByCodeGetStatesHandler(rw web.ResponseWriter, req *web.Request) {

	code := parseStringPathParameter(req, "code")
	var flag orm.Flag
	db.Where(&orm.Flag{Code: code}).First(&flag)

	if (flag.ID == 0) {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("Flag not found"))
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

	var flagReadings []orm.FlagState
	orm.GetOrderedWindowedQuery(db, "flag_id", flag.ID, start, end).Find(&flagReadings)
	marshal(rw, flagReadings)
}