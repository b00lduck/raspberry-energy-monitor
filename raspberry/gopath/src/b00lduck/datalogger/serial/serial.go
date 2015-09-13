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

func processDatagram(data []byte) error {

    fmt.Println(data)

    if data[37] != 10 {
	return errors.New("last char in datagram must be newline (0x0a)")
    }

    // Check ADC values
    for i:=0;i<8;i++ {
	index := 4 * i
	// check first ADC digit (0,1,2,3)
	if err := isSmallHexDigit(data,index); err != nil {
	    return err	
	}

	// check second and third ADC digit (0-9,a-f)
	for i:=0;i<2;i++ {
	    index += 1
	    if err := isHexDigit(data,index); err != nil {
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

    if err := isHexDigit(data,32); err != nil {
	return err
    }

    if err := isHexDigit(data,33); err != nil {
	return err
    }

    // check CRC value (format only, real CRC check later)
    if data[34] != 32 {
        return errors.New("char at index 34 must be a space (0x20)")
    }

    if err := isHexDigit(data,35); err != nil {
	return err
    }

    if err := isHexDigit(data,36); err != nil {
	return err
    }

    adc_brauchwasser := parseADCSensorC(5, data) 
    adc_aussen := parseADCSensorB(6, data) 
    adc_kessel := parseADCSensorA(7, data) 

    fmt.Println("Brauchwasser: " + fmt.Sprintf("%.1f", adc_brauchwasser) + " C")
    fmt.Println("Aussen: " + fmt.Sprintf("%.1f", adc_aussen) + " C")
    fmt.Println("Kessel: " + fmt.Sprintf("%.1f", adc_kessel) + " C")

    return nil
}

type interpol struct {
    t float32
    u float32
}

// Sensor type A
func parseADCSensorA(ch int, data []byte) (ret float32) {
    points := []interpol {
	interpol { 20, 2.47 },
	interpol { 25, 2.37 },
	interpol { 30, 2.27 },
	interpol { 35, 2.17 },
	interpol { 40, 2.06 },
	interpol { 45, 1.94 },
	interpol { 50, 1.82 },
	interpol { 55, 1.70 },
	interpol { 60, 1.57 },
	interpol { 65, 1.46 },
	interpol { 70, 1.34 },
	interpol { 75, 1.23 },
	interpol { 80, 1.13 },
    }
    volt := parseADCVolt(ch, data)
    return interpolate(volt, points)
}

func interpolate(u float32, table []interpol) float32 {

    if u >= table[0].u {
	return table[0].t
    }

    for index := 0; index < len(table); index++ {
	if (table[index].u < u) {
	    interval := table[index-1].u - table[index].u	
	    a := table[index-1].u - u
	    frac := a / interval
	    
	    t_interval := table[index-1].t - table[index].t
	    
	    return table[index-1].t - t_interval * frac
	}
    }

    return table[len(table)-1].t

}

// Sensor type B
func parseADCSensorB(ch int, data []byte) (ret float32) {
    points := []interpol {
	interpol { -20, 4.54 },
	interpol { -15, 4.42 },
	interpol { -10, 4.29 },
	interpol { -5, 4.13 },
	interpol { 0, 3.96 },
	interpol { 5, 3.77 },
	interpol { 10, 3.56 },
	interpol { 15, 3.34 },
	interpol { 20, 3.05 },
    }
    volt := parseADCVolt(ch, data)
    return interpolate(volt, points)
}

// Sensor type C
func parseADCSensorC(ch int, data []byte) (ret float32) {
    points := []interpol {
	interpol { 20, 2.60 },
	interpol { 25, 2.47 },
	interpol { 30, 2.34 },
	interpol { 35, 2.20 },
	interpol { 40, 2.06 },
	interpol { 45, 1.91 },
	interpol { 50, 1.77 },
	interpol { 55, 1.63 },
	interpol { 60, 1.49 },
	interpol { 65, 1.36 },
	interpol { 70, 1.23 },
	interpol { 75, 1.12 },
	interpol { 80, 1.01 },
    }
    volt := parseADCVolt(ch, data)
    return interpolate(volt, points)
}

func parseADCVolt(ch int, data []byte) (ret float32) {
    return float32(parseADC(ch, data)) * (4.97 / 1024.0)
}

func parseADC(ch int, data []byte) (ret uint16) {
    index := ch * 4
    ret = parseHexDigit(index, data) * 256
    ret += parseHexDigit(index + 1, data) * 16
    ret += parseHexDigit(index + 2, data)
    return
}

func parseHexDigit(index int, data []byte) uint16 {
    val := uint16(data[index])
    if val > 57 {
	return val - 87
    }
    return val - 48
}

func isSmallHexDigit(data []byte, index int) error {
    c := data[index]
    if c < 48 || c > 52 {
        return errors.New("char at index " + strconv.Itoa(index) + " must be a valid small hex digit (0-3) but was " + string(c))
    }
    return nil
}

func isHexDigit(data []byte, index int) error {
    c := data[index]
    if c < 48 || (c > 57 && (c < 97 || c > 102)) {
        return errors.New("char at index " + strconv.Itoa(index) + " must be a valid lowercase hex digit (0-9,a-f) but was " + string(c))
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
