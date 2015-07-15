#pragma once
#include <X11/Xlib.h>
#include <X11/Xutil.h>
#include <X11/Xos.h>
#include <X11/Xatom.h>

extern Display *display;
extern Window win;
extern Visual *visual;

void window_create();
void window_close();