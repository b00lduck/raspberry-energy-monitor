#pragma once
#include <X11/Xlib.h>

XImage* createImageFromFb(Display *display, Visual *visual);