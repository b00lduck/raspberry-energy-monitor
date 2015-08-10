package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"b00lduck/tools"
	"b00lduck/datalogger/dataservice/orm"
	"b00lduck/datalogger/dataservice/initialization"
	"b00lduck/datalogger/dataservice/rest"
	"time"
)

func main() {
	fmt.Println("START")

	db, err := gorm.Open("mysql", "root:root@/rem-dataservice?parseTime=true")
	tools.ErrorCheck(err)

	db.SingularTable(true)

	db.AutoMigrate(&orm.Counter{}, &orm.CounterEvent{})

	db.Model(&orm.CounterEvent{}).AddForeignKey("counter_id", "counter(id)", "RESTRICT", "RESTRICT")

	counterChecker := initialization.NewCounterChecker(&db)
	counterChecker.CheckCounters()

	go rest.StartServer(&db)

	for {
		time.Sleep(1 * time.Second)
	}
}


