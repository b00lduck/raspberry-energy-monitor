package main

import (
	"fmt"
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"b00lduck/tools"
)

type Counter struct {
	ID          int
	Name		string
	Unit 		string
	Reading		float32
	LastTick	time.Time
}

func main() {
	fmt.Println("START")

	db, err := gorm.Open("mysql", "root:root@/rem-dataservice")
	tools.ErrorCheck(err)

	// Get database connection handle [*sql.DB](http://golang.org/pkg/database/sql/#DB)
	db.DB()

	// Then you could invoke `*sql.DB`'s functions with it
	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	// Disable table name's pluralization
	db.SingularTable(true)
}

