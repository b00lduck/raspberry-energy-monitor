#include <stdbool.h>
#include <pthread.h>

#include "exposeEventThread.h"
#include "config.h"
#include "window.h"

pthread_t exposeEventThread;
bool exposeEventThreadStopRequest = false;

/**
 * Periodically generates an XExpose event to trigger the redraw of the screen.
 *
 * Runs in own thread until endExposeEventThread is true.
 */
void *exposeEventThreadHandler (void* q) {

    XEvent exposeEvent;
    memset(&exposeEvent, 0, sizeof(exposeEvent));
    exposeEvent.type = Expose;
    exposeEvent.xexpose.window = win;

     while (!exposeEventThreadStopRequest) {
        XSendEvent(display, win, False, ExposureMask, &exposeEvent);
        XFlush(display);
        usleep (1000000 / WINDOW_REFRESH_RATE);
    }

    return NULL;
}

void exposeEventThread_start() {
    pthread_create(&exposeEventThread, NULL, exposeEventThreadHandler, NULL);
}

void exposeEventThread_stop() {
    exposeEventThreadStopRequest = true;
    void* retval;
    pthread_join(exposeEventThread, &retval);
}