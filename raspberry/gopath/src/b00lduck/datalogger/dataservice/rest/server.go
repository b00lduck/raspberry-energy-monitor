package rest
import (
	"github.com/jinzhu/gorm"
	"net/http"
	"b00lduck/tools"
	"b00lduck/datalogger/dataservice/rest/context"
	"b00lduck/datalogger/dataservice/rest/handler"
)

func StartServer(db *gorm.DB) {

	context := context.NewContext(db)

	http.Handle("/counter", HttpHandler{context, handler.CounterHandler})

	e := http.ListenAndServe(":8080", nil)
	tools.ErrorCheck(e)
}
