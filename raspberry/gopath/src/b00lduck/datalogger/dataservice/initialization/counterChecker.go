package initialization
import (
	"github.com/jinzhu/gorm"
	"b00lduck/datalogger/dataservice/orm"
)

type CounterChecker struct {
	db *gorm.DB
}

var counters = []orm.Counter {
	{ Code:	"GAS_1", Name: "Erdgas", Unit: "m³"},
	{ Code:	"WAT_1", Name: "Wasser Hauptzähler", Unit: "m³"},
	{ Code:	"ELE_1", Name: "Strom", Unit: "kWh"}}

func NewCounterChecker(db *gorm.DB) (counterChecker *CounterChecker) {
	counterChecker = new(CounterChecker)
	counterChecker.db = db
	return
}

func (c *CounterChecker) CheckCounters() {
	for i := range counters {
		c.checkCounter(counters[i])
	}
}

// Check if the given counter exists in the database and create it if not
func (c *CounterChecker) checkCounter(counter orm.Counter) {
	c.db.Where(orm.Counter{Code: counter.Code}).FirstOrCreate(&counter)
}