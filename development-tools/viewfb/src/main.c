/**
 * main.c
 *
 * @author Daniel Zerlett (daniel@zerlett.eu)
 **/
#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <stdbool.h>

#include "config.h"
#include "framebuffer.h"
#include "touchscreen.h"
#include "window.h"
#include "exposeEventThread.h"
#include "imageConverter.h"

void cleanup() {
    exposeEventThread_stop();

    window_close();

    fb_uninit();
    ts_uninit();
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

    window_create();

    exposeEventThread_start(win);

    while(1) {

        XEvent report;
        XNextEvent(display, &report);

        switch ( report.type ) {

            case Expose:
                if ( report.xexpose.count != 0 ) break;
                XImage* ximage = createImageFromFb(display, visual);
                XPutImage(display, win, DefaultGC(display, 0), ximage, 0, 0, 0, 0, FRAMEBUFFER_XRES, FRAMEBUFFER_YRES);
                XDestroyImage(ximage);
                break;

            case ButtonPress:
                ts_press(report.xbutton.x, report.xbutton.y);
                break;

            case ButtonRelease:
                ts_release();
                break;

            case KeyPress:
                cleanup();
                exit(EXIT_SUCCESS);

        }

    }

    return 0;

}


