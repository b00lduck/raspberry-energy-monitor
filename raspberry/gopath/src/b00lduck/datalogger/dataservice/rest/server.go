package rest
import (
	"github.com/jinzhu/gorm"
	"net/http"
	"b00lduck/tools"
	"github.com/gocraft/web"
	"encoding/json"
	"log"
	"strconv"
	"net/url"
)

var db *gorm.DB

type Context struct {
	values *url.Values
}

func StartServer(database *gorm.DB) {

	db = database

	router := web.New(Context{}).
	Middleware(web.LoggerMiddleware).
	Middleware((*Context).QueryVarsMiddleware).
	Middleware(CorsMiddleware).

	Get ("/counter", 			(*Context).CounterHandler).
	Get ("/counter/:id", 		(*Context).CounterByIdHandler).
	Post("/counter/:id/tick", 	(*Context).CounterByIdTickHandler).
	Put ("/counter/:id/corr", 	(*Context).CounterByIdCorrectHandler).
	Get ("/counter/:id/events",	(*Context).CounterByIdEventsHandler)

	e := http.ListenAndServe(":8080", router)
	tools.ErrorCheck(e)
}

func (c *Context) QueryVarsMiddleware(rw web.ResponseWriter, r *web.Request, next web.NextMiddlewareFunc) {

	values, err := parseQueryParams(rw, r)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Malformed URL"))
		return
	}

	c.values = &values
	next(rw, r)
}

func CorsMiddleware(rw web.ResponseWriter, r *web.Request, next web.NextMiddlewareFunc) {

	rw.Header().Set("Access-Control-Allow-Origin", "*")

	next(rw, r)
}

func marshal(rw web.ResponseWriter, v interface{}) {

	buffer, err := json.Marshal(v)
	if err != nil {
		log.Println("Error marshalling JSON")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")

	if _, err = rw.Write(buffer); err != nil {
		log.Println("Error writing JSON to buffer")
		rw.WriteHeader(http.StatusInternalServerError)
	}

}

func parseUintFromString(s string) (ret uint64, err error) {
	ret, err = strconv.ParseUint(s, 10, 64)
	return
}

func (c *Context) parseUintQueryParameter(rw web.ResponseWriter, name string) (ret uint64, err error) {

	s := c.values.Get(name)
	if s == "" {
		return 0, nil
	}

	ret,err = parseUintFromString(s)
	if err != nil {
		log.Println("Error parsing uint64 " + name +": " + s)
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Malformed parameter " + name))
	}

	return
}

func parseUintPathParameter(rw web.ResponseWriter, req *web.Request, name string) (id uint64, err error) {

	s := req.PathParams[name]
	id,err = parseUintFromString(s)

	if (err != nil) {
		log.Println("Error parsing uint64 " + name +": " + s)
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Malformed path parameter"))
	}

	return
}

func parseQueryParams(rw web.ResponseWriter, req *web.Request) (values url.Values, err error) {

	u,err := url.Parse(req.RequestURI)
	if err != nil {
		return
	}

	values,err = url.ParseQuery(u.RawQuery)

	return
}
