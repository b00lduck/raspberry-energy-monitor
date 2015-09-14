package initialization
import (
	"github.com/jinzhu/gorm"
	"b00lduck/datalogger/dataservice/orm"
)

type CounterChecker struct {
	db *gorm.DB
}

var counters = []orm.Counter {
	{ Code:	"GAS_1", Name: "Erdgas", Unit: "m³", TickAmount: 100},
	{ Code:	"WAT_1", Name: "Wasser Hauptzähler", Unit: "m³", TickAmount: 10},
	{ Code:	"ELE_1", Name: "Strom", Unit: "kWh", TickAmount: 10}}

var thermometers = []orm.Thermometer {
	{ Code:	"HEIZ_AUSSEN", Name: "Heizung Aussentemperatur"},
	{ Code:	"HEIZ_KESSEL", Name: "Heizung Kesseltemperatur"},
	{ Code:	"HEIZ_BRAUCHW", Name: "Heizung Brauchwassertemperatur"},
}

func NewCounterChecker(db *gorm.DB) (counterChecker *CounterChecker) {
	counterChecker = new(CounterChecker)
	counterChecker.db = db
	return
}

func (c *CounterChecker) CheckCounters() {
	for i := range counters {
		c.checkCounter(counters[i])
	}
	for i := range thermometers {
		c.checkThermometer(thermometers[i])
	}
}

// Check if the given counter exists in the database and create it if not
func (c *CounterChecker) checkCounter(counter orm.Counter) {
	c.db.Where(orm.Counter{Code: counter.Code}).FirstOrCreate(&counter)
}

// Check if the given thermometer exists in the database and create it if not
func (c *CounterChecker) checkThermometer(thermometer orm.Thermometer) {
	c.db.Where(orm.Thermometer{Code: thermometer.Code}).FirstOrCreate(&thermometer)
}