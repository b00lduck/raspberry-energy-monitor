/**
 *
 * viewfb.c
 *
 * Display contents of virtual framebuffer device in x window
 *
 * @author Daniel Zerlett
 *
 **/

#include <X11/Xlib.h>
#include <X11/Xutil.h>
#include <X11/Xos.h>
#include <X11/Xatom.h>
#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <stdbool.h>
#include <pthread.h>

#include "framebuffer.h"
#include "touchscreen.h"
#include "config.h"

Display *display;
Window win;
static char *appname;

bool endExposeEventThread = false;

/**
 * Periodically generates an XExpose event to trigger the redraw of the screen.
 *
 * Runs in own thread until endExposeEventThread is true.
 */
void *tp (void* q) {
    XEvent exposeEvent;
    memset(&exposeEvent, 0, sizeof(exposeEvent));
    exposeEvent.type = Expose;
    exposeEvent.xexpose.window = win;

    while (!endExposeEventThread) {
        XSendEvent(display, win, False, ExposureMask, &exposeEvent);
        XFlush(display);
        usleep (1000000 / WINDOW_REFRESH_RATE);
    }

    return NULL;
}

/**
 * Create an X11 image from the pixel data in the framebuffer bitmap.
 * Do an B5-G6-R5 color conversion to R16-G16-B16
 */
XImage* createImageFromFb(Display *display, Visual *visual) {

    short int* fb = fb_getBitmap();

    char *image32 = malloc(FRAMEBUFFER_PIXEL * 4);
    char *p = image32;

    unsigned int counter = FRAMEBUFFER_PIXEL;

    while (counter--) {
        *p++ = (*fb & 0b0000000000011111) << 3;
        *p++ = (*fb & 0b0000011111100000) >> 3;
        *p++ = (*fb++ & 0b1111100000000000) >> 8;
        p++;
    }
    return XCreateImage(display, visual, 24, ZPixmap, 0, image32, FRAMEBUFFER_XRES, FRAMEBUFFER_YRES, 32, 0);
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

    printf("creating window\n");

    char *window_name = "Viewing virtual framebuffer";
    char *icon_name = "viewfb";

    char *display_name = NULL;

    XSizeHints *size_hints;
    XWMHints *wm_hints;
    XClassHint *class_hints;
    XTextProperty windowName, iconName;
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
    int x = DisplayWidth(display, screen_num) - FRAMEBUFFER_XRES;
    int y = 30;

    win = XCreateSimpleWindow(display, RootWindow(display, screen_num), x, y, FRAMEBUFFER_XRES, FRAMEBUFFER_YRES, 0,
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
    size_hints->min_width = FRAMEBUFFER_XRES;
    size_hints->min_height = FRAMEBUFFER_YRES;

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

    printf("entering event loop");

    while ( 1 ) {
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
                // cleanup and close
                endExposeEventThread = true;
                void* retval;
                pthread_join(thread, &retval);
                XFreeGC(display, gc);
                XCloseDisplay(display);
                fb_uninit();
                ts_uninit();
                exit(EXIT_SUCCESS);
 
        }

    }

}
