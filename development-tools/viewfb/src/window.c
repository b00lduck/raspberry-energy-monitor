#include <X11/Xlib.h>
#include <X11/Xutil.h>
#include <X11/Xos.h>
#include <X11/Xatom.h>
#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <stdbool.h>
#include <pthread.h>

#include "config.h"
#include "exposeEventThread.h"

Display *display;
Window win;
Visual *visual;
GC gc;

void window_close() {
    XFreeGC(display, gc);
    XCloseDisplay(display);
}

void window_create() {

    printf("Creating window...\n");

    char *window_name = "Viewing virtual framebuffer";
    char *icon_name = "viewfb";
    char *appname = "viewfb";
    char *display_name = NULL;

    XSizeHints *size_hints;
    XWMHints *wm_hints;
    XClassHint *class_hints;
    XTextProperty windowName, iconName;
    XGCValues values;

    if ( !(size_hints = XAllocSizeHints()) || !(wm_hints = XAllocWMHints()) || !(class_hints = XAllocClassHint())) {
        fprintf(stderr, "%s: couldn't allocate memory.\n", appname);
        exit(EXIT_FAILURE);
    }

    if ( (display = XOpenDisplay(display_name)) == NULL ) {
        fprintf(stderr, "%s: couldn't connect to X server %s\n", appname, display_name);
        exit(EXIT_FAILURE);
    }

    int screen_num = DefaultScreen(display);
    int x = DisplayWidth(display, screen_num) - WINDOW_XRES;
    int y = 30;

    win = XCreateSimpleWindow(display, RootWindow(display, screen_num), x, y, WINDOW_XRES, WINDOW_YRES, 0,
        BlackPixel(display, screen_num), WhitePixel(display, screen_num));

    if ( XStringListToTextProperty(&window_name, 1, &windowName) == 0 ) {
        fprintf(stderr, "%s: structure allocation for windowName failed.\n", appname);
        exit(EXIT_FAILURE);
    }

    if ( XStringListToTextProperty(&icon_name, 1, &iconName) == 0 ) {
        fprintf(stderr, "%s: structure allocation for iconName failed.\n", appname);
        exit(EXIT_FAILURE);
    }

    size_hints->flags  = PPosition | PSize | PMinSize;
    size_hints->min_width = WINDOW_XRES;
    size_hints->min_height = WINDOW_YRES;

    wm_hints->flags         = StateHint | InputHint;
    wm_hints->initial_state = NormalState;
    wm_hints->input         = True;

    class_hints->res_name   = appname;
    class_hints->res_class  = "viewfb";

    XSetWMProperties(display, win, &windowName, &iconName, NULL, 0, size_hints, wm_hints, class_hints);

    XSelectInput(display, win, ExposureMask | KeyPressMask | ButtonPressMask | ButtonReleaseMask );

    gc = XCreateGC(display, win, 0, &values);

    visual = DefaultVisual(display, 0);

    XMapWindow(display, win);

}

