package orm
import (
)

type EventType uint8

const (
	TICK EventType = 1 + iota
	ABS_CORR
	ABS_READ
)

type CounterEvent struct {
	ID        	uint 	`gorm:"primary_key"`
	CounterID	uint 	`sql:"index"`
	Timestamp	int64 	`sql:"index`
	EventType 	EventType
	Delta		uint32  // delta multiplied by 1000
	Reading		uint64  // reading multiplied by 1000
}