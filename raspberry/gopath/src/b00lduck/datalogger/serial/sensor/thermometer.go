package sensor
import (
	"b00lduck/tools"
	"math"
	"b00lduck/datalogger/dataservice/client"
	"fmt"
)

type Thermometer struct {
	oldValue float64
	precision float64
	code string
}

func NewThermometer(code string, precision float64) Thermometer {
	return Thermometer{
		oldValue: 0,
		precision: precision,
		code: code,
	}
}

func (t *Thermometer) SetNewReading(reading float64) {

	// precision reduction
	limitedPrecisionValue := tools.Round(reading / t.precision) * t.precision

	if math.Abs(float64(limitedPrecisionValue - t.oldValue)) > t.precision {
		if err := client.SendThermometerReading(t.code, limitedPrecisionValue); err != nil {
			fmt.Println(err)
		} else {
			t.oldValue = limitedPrecisionValue
		}
	}
}

