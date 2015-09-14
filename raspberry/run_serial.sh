#!/bin/bash

export GOROOT=/home/pi/go
export GOPATH=/home/pi/raspberry-energy-monitor/raspberry/gopath

echo "GOPATH is $GOPATH"
echo "GOROOT is $GOROOT"

cd $GOPATH

/home/pi/go/bin/go get ./...

/home/pi/go/bin/go run $GOPATH/src/b00lduck/datalogger/serial/serial.go