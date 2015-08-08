package context
import (
	"github.com/jinzhu/gorm"
)

type Context struct {
	DB *gorm.DB
}

func NewContext(db *gorm.DB) *Context {
	newContext := new(Context)
	newContext.DB = db
	return newContext
}

