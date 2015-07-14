# raspberry-energy-monitor
A residential energy consumption logger with the Raspberry PI 2

# Requirements

## Raspberry 
### Hardware
 - Raspberry PI 2
 - SD-Card with Raspbian image
 - RPi 2.8" display with touchscreen from Watterott
 - 3-Axis magnetometer breakout board

### Operating system and drivers
 - Raspbian installation
 - driver for display configured as /dev/fb1
 - Driver for touchscreen with ADS7846 configured as /dev/input/event0
 - shell and x11 disabled on /dev/fb1
 - fbcpd disabled
 
### Software
 - go
 
## Development machine
 - Linux kernel with following extra features:
  - vfb kernel module 
   - start with modprobe vfb vfb_enable=yes
  - uinput kernel module
 - 
