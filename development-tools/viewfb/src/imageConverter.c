#include <stdlib.h>
#include <string.h>

#include "config.h"
#include "imageConverter.h"
#include "framebuffer.h"

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