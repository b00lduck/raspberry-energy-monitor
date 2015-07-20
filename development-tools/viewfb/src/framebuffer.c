#include <unistd.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <signal.h>
#include <fcntl.h>
#include <linux/fb.h>
#include <sys/mman.h>
#include <sys/ioctl.h>
#include <sys/stat.h>
#include "config.h"

short int *fb_p = 0;
int fb_fd = 0;
struct fb_var_screeninfo fb_old_vinfo;


int fb_setResolution() {
	struct fb_var_screeninfo fb_vinfo;

	// Get variable screen information
	if (ioctl(fb_fd, FBIOGET_VSCREENINFO, &fb_vinfo)) {
		printf("Error reading variable information from framebuffer device.\n");
		return 1;
	}
	printf("Original resolution %dx%d, %dbpp\n", fb_vinfo.xres, fb_vinfo.yres, fb_vinfo.bits_per_pixel );

	fb_vinfo.xres = FRAMEBUFFER_XRES;
	fb_vinfo.yres = FRAMEBUFFER_YRES;
	fb_vinfo.bits_per_pixel = FRAMEBUFFER_BITS_PER_PIXEL;
	fb_vinfo.xoffset = 0;
	fb_vinfo.xres_virtual = 0;
	fb_vinfo.yres_virtual = 0;
	fb_vinfo.pixclock = 260415;
	fb_vinfo.left_margin = 0;
	fb_vinfo.right_margin = 0;
	fb_vinfo.upper_margin = 0;
	fb_vinfo.lower_margin = 0;
	fb_vinfo.hsync_len = 0;
	fb_vinfo.vsync_len = 0;

	if (ioctl(fb_fd, FBIOPUT_VSCREENINFO, &fb_vinfo)) {
		printf("Error setting variable information of framebuffer device.\n");
		return 1;
	}

	// Get variable screen information
	if (ioctl(fb_fd, FBIOGET_VSCREENINFO, &fb_vinfo)) {
		printf("Error reading variable information from framebuffer device.\n");
		return 1;
	}

    printf("Current resolution: %dx%d, %dbpp\n", fb_vinfo.xres, fb_vinfo.yres, fb_vinfo.bits_per_pixel );

	if ((fb_vinfo.xres != FRAMEBUFFER_XRES)
	    || (fb_vinfo.xres != FRAMEBUFFER_YRES)
	    || (fb_vinfo.bits_per_pixel != FRAMEBUFFER_BITS_PER_PIXEL)) {
		printf("Error setting new resolution of framebuffer device.\n");
    	return 1;
	}

	return 0;

}

void fb_uninit() {
    munmap(fb_p, FRAMEBUFFER_BYTES);
    close(fb_fd);
}

int fb_init(char *device) {

	// Open the file for reading and writing
	fb_fd = open(device, O_RDWR);
	if (!fb_fd) {
	    printf("Error: cannot open framebuffer device %s.\n", device);
	    return 1;
	}
	printf("The framebuffer device %s was opened successfully.\n", device);

	if (fb_setResolution()) {
	    printf("Error: cannot set framebuffer parameters of %s.\n", device);
	};

	// map fb to user mem 
	fb_p = mmap(0, FRAMEBUFFER_BYTES, PROT_READ | PROT_WRITE, MAP_SHARED, fb_fd, 0);

	if ((long)fb_p == -1) {
		printf("Failed to mmap %s.\n", device);
		return 1;
	}

    int i = strtol("0666", 0, 8);
    if (chmod (device, i) < 0) {
        printf("Error: cannot chmod %s to 666", device);
        return 1;
    }

	return 0;
}

short int* fb_getBitmap() {
    return fb_p;
}
