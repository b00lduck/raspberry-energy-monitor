package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"b00lduck/tools"
	"b00lduck/datalogger/dataservice/orm"
	"b00lduck/datalogger/dataservice/initialization"
	"b00lduck/datalogger/dataservice/rest"
)

func main() {
	db, err := gorm.Open("mysql", "root:root@/rem-dataservice?parseTime=true")
	tools.ErrorCheck(err)

	db.SingularTable(true)

	db.AutoMigrate(&orm.Counter{}, &orm.CounterEvent{}, &orm.Thermometer{}, &orm.ThermometerReading{})

	db.Model(&orm.CounterEvent{}).AddForeignKey("counter_id", "counter(id)", "RESTRICT", "RESTRICT")
	db.Model(&orm.ThermometerReading{}).AddForeignKey("thermometer_id", "thermometer(id)", "RESTRICT", "RESTRICT")

	counterChecker := initialization.NewCounterChecker(&db)
	counterChecker.CheckCounters()

	rest.StartServer(&db)
}


