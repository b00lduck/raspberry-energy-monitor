package rest
import (
	"github.com/jinzhu/gorm"
	"net/http"
	"b00lduck/tools"
	"github.com/gocraft/web"
)

var db *gorm.DB

type Context struct {
}

func StartServer(database *gorm.DB) {

	db = database

	router := web.New(Context{}).
	Middleware(web.LoggerMiddleware).
	//Middleware(web.ShowErrorsMiddleware).

	Get("/counter", (*Context).CounterHandler).
	Get("/counter/:id", (*Context).CounterHandlerId)

	e := http.ListenAndServe(":8080", router)
	tools.ErrorCheck(e)
}
