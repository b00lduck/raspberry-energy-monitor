package handler

import (
	"net/http"
	"b00lduck/datalogger/dataservice/orm"
	"encoding/json"
	"b00lduck/tools"
	"b00lduck/datalogger/dataservice/rest/context"
)

func CounterHandler(ws *context.Context, writer http.ResponseWriter, r *http.Request) (int, error) {

	var counters []orm.Counter

	ws.DB.Find(&counters)

	buffer, err := json.Marshal(counters)
	tools.ErrorCheck(err)

	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err = writer.Write(buffer)
	tools.ErrorCheck(err)


	return http.StatusOK, nil
}
