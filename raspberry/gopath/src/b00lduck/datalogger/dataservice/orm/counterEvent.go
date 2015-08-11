package orm
import (
)

type CounterEvent struct {
	ID        	uint 	`gorm:"primary_key"`
	CounterID	uint 	`sql:"index"`
	Timestamp	int64 	`sql:"index`
	EventType 	uint8
	Delta		uint32  // delta multiplied by 1000
	Reading		uint64  // reading multiplied by 1000
}