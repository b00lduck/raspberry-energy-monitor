#!/bin/bash

screen -dmS dataservice ./run_dataservice.sh
screen -dmS serial ./run_serial.sh
screen -dmS display ./run_display.sh
screen -dmS acquisition ./run_acquisition.sh


