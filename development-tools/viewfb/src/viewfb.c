/**
 *
 * viewfb.c
 *
 * Display contents of virtual framebuffer device in x window
 *
 * copyleft Daniel Zerlett 2015
 * daniel@zerlett.eu
 *
 **/

#include <X11/Xlib.h>
#include <X11/Xutil.h>
#include <X11/Xos.h>
#include <X11/Xatom.h>
#include <stdlib.h>
#include <string.h>
#include <stdio.h>

#include "framebuffer.h"
#include "touchscreen.h"

Display *display;
Window win;
static char *appname;

void* tp (void* q) {
    XEvent exppp;
    while (1) {
        usleep (25000);
        memset(&exppp, 0, sizeof(exppp));
        exppp.type = Expose;
        exppp.xexpose.window = win;
        XSendEvent(display,win,False,ExposureMask,&exppp);
        XFlush(display);
    }
    return NULL;
}

XImage* createImageFromFb(Display *display, Visual *visual) {

    const unsigned int width = 320;
    const unsigned int height = 240;

    short int* fb = fb_getBitmap();

    int i, j;
    unsigned char* image32 = (unsigned char *) malloc(320*240*4);
    unsigned char* p = image32;
    for(i = 0; i < width; i++) {
        for(j=0; j < height; j++) {
            *p++ = (*fb++ & 0b0000000000011111) << 3;
            *p++ = (*fb & 0b0000011111100000) >> 3;
            *p++ = (*fb & 0b1111100000000000) >> 8;
            p++;
        }
    }
    return XCreateImage(display, visual, 24, ZPixmap, 0, image32, width, height, 32, 0);
}

int main(int argc, char * argv[]) {

    int rc = fb_init("/dev/fb2");

    if (rc != 0) {
        exit(EXIT_FAILURE);
    }

    rc = ts_init("/dev/uinput");

    if (rc != 0) {
        exit(EXIT_FAILURE);
    }

    unsigned int border_width = 0;
    char *window_name = "Viewing virtual framebuffer";
    char *icon_name = "viewfb";

    char *display_name = NULL;

    XSizeHints *size_hints;
    XWMHints *wm_hints;
    XClassHint *class_hints;
    XTextProperty windowName, iconName;
    XEvent report;
    XGCValues values;

    appname = argv[0];

    if ( !(size_hints = XAllocSizeHints()) || !(wm_hints = XAllocWMHints()) || !(class_hints = XAllocClassHint())) {
        fprintf(stderr, "%s: couldn't allocate memory.\n", appname);
        exit(EXIT_FAILURE);
    }

    if ( (display = XOpenDisplay(display_name)) == NULL ) {
        fprintf(stderr, "%s: couldn't connect to X server %s\n", appname, display_name);
        exit(EXIT_FAILURE);
    }

    int screen_num = DefaultScreen(display);
    unsigned int width = 320;
    unsigned int height = 240;
    int x = DisplayWidth(display, screen_num) - width;
    int y = 30;

    win = XCreateSimpleWindow(display, RootWindow(display, screen_num), x, y, width, height, border_width, 
        BlackPixel(display, screen_num), WhitePixel(display, screen_num));

    if ( XStringListToTextProperty(&window_name, 1, &windowName) == 0 ) {
        fprintf(stderr, "%s: structure allocation for windowName failed.\n", appname);
        exit(EXIT_FAILURE);
    }

    if ( XStringListToTextProperty(&icon_name, 1, &iconName) == 0 ) {
        fprintf(stderr, "%s: structure allocation for iconName failed.\n", appname);
        exit(EXIT_FAILURE);
    }

    size_hints->flags       = PPosition | PSize | PMinSize;
    size_hints->min_width   = 320;
    size_hints->min_height  = 240;

    wm_hints->flags         = StateHint | InputHint;
    wm_hints->initial_state = NormalState;
    wm_hints->input         = True;

    class_hints->res_name   = appname;
    class_hints->res_class  = "viewfb";

    XSetWMProperties(display, win, &windowName, &iconName, argv, argc, size_hints, wm_hints, class_hints);

    XSelectInput(display, win, ExposureMask | KeyPressMask | ButtonPressMask | ButtonReleaseMask );

    GC gc = XCreateGC(display, win, 0, &values);

    Visual *visual = DefaultVisual(display, 0);

    XMapWindow(display, win);

    pthread_t thread;
    pthread_create(&thread, NULL, tp, NULL);

    while ( 1 ) {
        XEvent report;
        XNextEvent(display, &report);

        switch ( report.type ) {

            case Expose:
                if ( report.xexpose.count != 0 ) break;
                XImage* ximage = createImageFromFb(display, visual);
                XPutImage(display, win, DefaultGC(display, 0), ximage, 0, 0, 0, 0, 320, 240);
                break;

            case ButtonPress:
                ts_press(report.xbutton.x, report.xbutton.y);
                break;

            case ButtonRelease:
                ts_release();
                break;

            case KeyPress:
                XFreeGC(display, gc);
                XCloseDisplay(display);
                fb_uninit();
                exit(EXIT_SUCCESS);
 
        }

    }

}
