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
    if math.Abs(float64(adc_brauchwasser - oldval_brauchwasser)) > 0.2 {
		fmt.Println("Brauchwasser: " + fmt.Sprintf("%.1f", adc_brauchwasser) + " C")
		oldval_brauchwasser = adc_brauchwasser
	}

	adc_aussen := parser.ParseADCSensorB(6, data)
	if math.Abs(float64(adc_aussen - oldval_aussen)) > 0.2 {
		fmt.Println("Aussen: " + fmt.Sprintf("%.1f", adc_aussen) + " C")
		oldval_aussen = adc_aussen
	}

    adc_kessel := parser.ParseADCSensorA(7, data)
	if math.Abs(float64(adc_kessel - oldval_kessel)) > 0.2 {
		fmt.Println("Kessel: " + fmt.Sprintf("%.1f", adc_kessel) + " C")
		oldval_kessel = adc_kessel
	}

    return nil
}

func sendTick() {

	client := &http.Client{}
	request, err := http.NewRequest("POST", "http://localhost:8080/counter/1/tick", strings.NewReader(""))
	if err != nil {
		fmt.Println("Error creating tick request to dataservice")
		fmt.Println(err)
		return
	}
	request.ContentLength = 0
	_, err = client.Do(request)
	if err != nil {
		fmt.Println("Error sending tick request to dataservice")
		fmt.Println(err)
	}

}
