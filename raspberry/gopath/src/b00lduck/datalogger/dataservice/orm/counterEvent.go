package orm
import (
	"time"
)

type CounterEvent struct {
	ID        	uint `gorm:"primary_key"`
	Timestamp	time.Time
	EventType 	uint8
	Delta		uint32
	Reading		uint64
}