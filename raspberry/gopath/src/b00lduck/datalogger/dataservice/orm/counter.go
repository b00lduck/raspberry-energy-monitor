package orm
import (
	"github.com/jinzhu/gorm"
	"time"
)

type Counter struct {
	gorm.Model
	Code		string
	Name		string
	Unit 		string
	Reading		float32
	LastTick	time.Time
}