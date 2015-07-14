# viewfb - A X11 Framebuffer viewer with touchscreen emulator

## Description

This tool shows the contents of a framebuffer device in a X11 window and translates klicks to touchscreen events.

## Usage

 - First parameter: vfb framebuffer device
 - Second parameter: uinput device for touchscreen mocking

./viewfb /dev/fb2 /dev/input/event11