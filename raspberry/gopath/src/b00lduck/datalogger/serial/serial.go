package main

import (
	"fmt"
	"time"
	"net/http"
	"strings"
	"github.com/tarm/serial"
	"log"
	"errors"
	"strconv"
	"math"
	"b00lduck/datalogger/serial/parser"
	"io/ioutil"
)

func main() {

    c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600, ReadTimeout: time.Second * 2}
    s, err := serial.OpenPort(c)

    if err != nil {
        log.Fatal(err)
    }
    
    for {
		err = requestDatagram(s)
		if err != nil {
			fmt.Println(err)
			time.Sleep(1 * time.Second)
		} else {
			time.Sleep(5 * time.Second)
		}
    }
}

func requestDatagram(s *serial.Port) error {
    
    n, err := s.Write([]byte("DUMP\r"))
    if err != nil {
        return err
    }

    buf := make([]byte, 38) // Datagram length is 38 bytes incl. \n
    n, err = s.Read(buf)
    if err != nil {
        return err
    }

    // check for correct size
    if n != 38 {
	return errors.New("received datagram with invalid size (must: 38, was: " + strconv.Itoa(n) + ")")
    }

    return processDatagram(buf)

}

var oldval_brauchwasser = float32(0)
var oldval_aussen = float32(0)
var oldval_kessel = float32(0)

func processDatagram(data []byte) error {

    if data[37] != 10 {
	return errors.New("last char in datagram must be newline (0x0a)")
    }

    // Check ADC values
    for i:=0;i<8;i++ {
	index := 4 * i
	// check first ADC digit (0,1,2,3)
	if err := parser.IsSmallHexDigit(data,index); err != nil {
	    return err	
	}

	// check second and third ADC digit (0-9,a-f)
	for i:=0;i<2;i++ {
	    index += 1
	    if err := parser.IsHexDigit(data,index); err != nil {
		return err
	    }
	}

	// check spaces between values
	index += 1
        if data[index] != 32 {
    	    return errors.New("char at index " + strconv.Itoa(index) + " must be a space (0x20)")
	}    
    }

    // check DIGITAL values
    if data[31] != 32 {
        return errors.New("char at index 31 must be a space (0x20)")
    }

    if err := parser.IsHexDigit(data,32); err != nil {
	return err
    }

    if err := parser.IsHexDigit(data,33); err != nil {
	return err
    }

    // check CRC value (format only, real CRC check later)
    if data[34] != 32 {
        return errors.New("char at index 34 must be a space (0x20)")
    }

    if err := parser.IsHexDigit(data,35); err != nil {
	return err
    }

    if err := parser.IsHexDigit(data,36); err != nil {
	return err
    }

    adc_brauchwasser := parser.ParseADCSensorC(5, data)
	adc_brauchwasser = Round(adc_brauchwasser / 0.25) * 0.25 // precision reduction
    if math.Abs(float64(adc_brauchwasser - oldval_brauchwasser)) > 0.25 {
		sendReading("HEIZ_BRAUCHW", adc_brauchwasser)
		oldval_brauchwasser = adc_brauchwasser
	}

	adc_aussen := parser.ParseADCSensorB(6, data)
	adc_aussen  = Round(adc_aussen  / 0.5) * 0.5 // precision reduction
	if math.Abs(float64(adc_aussen - oldval_aussen)) > 0.5 {
		sendReading("HEIZ_AUSSEN", adc_aussen)
		oldval_aussen = adc_aussen
	}

    adc_kessel := parser.ParseADCSensorA(7, data)
	adc_kessel  = Round(adc_kessel  / 0.5) * 0.5 // precision reduction
	if math.Abs(float64(adc_kessel - oldval_kessel)) > 0.5 {
		sendReading("HEIZ_KESSEL", adc_kessel)
		oldval_kessel = adc_kessel
	}

    return nil
}

func Round(f float32) float32 {
	return float32(math.Floor(float64(f + .5)))
}

func sendReading(code string, temp float32) {

	fmt.Println(code + ": " + fmt.Sprintf("%.2f", temp) + " C")

	intval := fmt.Sprintf("%d", uint64(Round(temp * 1000)))

	client := &http.Client{}
	request, err := http.NewRequest("POST", "http://localhost:8080/thermometer/" + code + "/reading", strings.NewReader(intval))
	if err != nil {
		fmt.Println("Error creating thermometer request to dataservice")
		fmt.Println(err)
		return
	}
	request.ContentLength = 0
	x, err := client.Do(request)
	if err != nil {
		fmt.Println("Error sending thermometer request to dataservice")
		fmt.Println(err)
	}

	if x.StatusCode != 200 {
		fmt.Println("Error sending thermometer request to dataservice")
		str, _ := ioutil.ReadAll(x.Body)
		fmt.Println( string(str) )
	}

}
