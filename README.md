# raspberry-energy-monitor
A residential energy consumption logger with the Raspberry PI 2

# Requirements

## Raspberry PI 2
### Hardware
 - Raspberry PI 2
 - RPi 2.8" display with touchscreen from Watterott
 - GPIO-Adapter for display
 - SD-Card with Raspbian image

### Operating system and drivers
 - Raspbian 7 installation
 - Driver for display configured as /dev/fb1
 - Driver for touchscreen with ADS7846 configured as /dev/input/event0
 - Shell and x11 disabled on /dev/fb1
 - fbcpd disabled
 
### Software
 - go in /home/pi/go
 - mysql 5.5
 - gradle
 - nodejs 
   For installation of nodejs see https://github.com/joyent/node/wiki/installing-node.js-via-package-manager
 - grunt-cli
   sudo npm install -g grunt-cli
 
## Development machine
 - Linux kernel with following extra features:
  - vfb kernel module 
   - start with modprobe vfb vfb_enable=yes
  - uinput kernel module

# Architecture

 - Spring-Based dataservice for RESTful layer above the mySQL database with embedded tomcat
 - AngularJS based frontend for graphing and remote control 
 - go based display and touchscreen daemon
 - go based data acquisition daemon
 - apache as http proxy server

