#!/bin/bash

export GOROOT=/home/pi/go
export GOPATH=/home/pi/raspberry-energy-monitor/raspberry/display

echo "GOPATH is $GOPATH"
echo "GOROOT is $GOROOT"

cd $GOPATH

/home/pi/go/bin/go run $GOPATH/src/b00lduck/datalogger/display/main.go /dev/fb1 /dev/input/touchscreen &



