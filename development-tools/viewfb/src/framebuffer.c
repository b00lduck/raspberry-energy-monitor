#include <unistd.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <signal.h>
#include <fcntl.h>
#include <linux/fb.h>
#include <sys/mman.h>
#include <sys/ioctl.h>

// 'global' variables to store screen info
short int *fb_p = 0;
struct fb_var_screeninfo fb_vinfo;
struct fb_fix_screeninfo fb_finfo;
long int fb_screensize = 0;
int fb_fd = 0;

void fb_uninit() {
    munmap(fb_p, fb_screensize);
    close(fb_fd);
}

int fb_init(char *device) {

	// Open the file for reading and writing
	fb_fd = open(device, O_RDWR);
	if (!fb_fd) {
	    printf("Error: cannot open framebuffer device.\n");
	    return 1;
	}
	printf("The framebuffer device was opened successfully.\n");

	// Get variable screen information
	if (ioctl(fb_fd, FBIOGET_VSCREENINFO, &fb_vinfo)) {
		printf("Error reading variable information.\n");
		return 1;
	}
	printf("Original %dx%d, %dbpp\n", fb_vinfo.xres, fb_vinfo.yres, fb_vinfo.bits_per_pixel );

	// map fb to user mem 
	fb_screensize = fb_vinfo.xres * fb_vinfo.yres * 2;
	fb_p = mmap(0, fb_screensize, PROT_READ | PROT_WRITE, MAP_SHARED, fb_fd, 0);

	if ((long)fb_p == -1) {
		printf("Failed to mmap.\n");
		return 1;
	}

	return 0;
}

short int* fb_getBitmap() {
    return fb_p;
}
