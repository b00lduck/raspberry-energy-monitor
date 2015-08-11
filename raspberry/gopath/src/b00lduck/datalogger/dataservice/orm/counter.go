package orm
import (
	"github.com/jinzhu/gorm"
)

type Counter struct {
	gorm.Model
	Code			string			`sql:"index"`
	Name			string			`sql:"index"`
	Unit 			string
	Reading			uint64			// Reading multiplied by 1000
	LastTick		int64			// Timestamp of last tick
}