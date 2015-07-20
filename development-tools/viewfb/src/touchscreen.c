#include <unistd.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <fcntl.h>
#include <linux/input.h>
#include <sys/ioctl.h>
#include <sys/stat.h>
#include <linux/uinput.h>
#include <errno.h>

// 'global' variables to store touchscreen info
int ts_fd = 0;

void ts_uninit() {
    close(ts_fd);
}

int ts_init(char *device) {
    int ret;

	// Open the file for writing
	ts_fd = open(device, O_WRONLY | O_NONBLOCK);
	if (!ts_fd) {
	    printf("Error: cannot open touchscreen device.\n");
	    return 1;
	}
	printf("The touchscreen device was opened successfully.\n");

    ret = ioctl(ts_fd, UI_SET_EVBIT, EV_KEY);
    if (ret != 0) {
    	printf("Error setting EV_KEY.\n");
        perror(0);
        return 1;
    }

    ret = ioctl(ts_fd, UI_SET_KEYBIT, BTN_TOUCH);
    if (ret != 0) {
    	printf("Error setting BTN_TOUCH.\n");
        perror(0);
        return 1;
    }

    ret = ioctl(ts_fd, UI_SET_EVBIT, EV_SYN);
    if (ret != 0) {
    	printf("Error setting EV_SYN.\n");
        perror(0);
        return 1;
    }

    ret = ioctl(ts_fd, UI_SET_EVBIT, EV_ABS);
    if (ret != 0) {
    	printf("Error setting EV_ABS.\n");
        perror(0);
        return 1;
    }

    ret = ioctl(ts_fd, UI_SET_ABSBIT, ABS_X);
    if (ret != 0) {
    	printf("Error setting ABS_X.\n");
        perror(0);
        return 1;
    }

    ret = ioctl(ts_fd, UI_SET_ABSBIT, ABS_Y);
    if (ret != 0) {
    	printf("Error setting ABS_Y.\n");
        perror(0);
        return 1;
    }

    ret = ioctl(ts_fd, UI_SET_ABSBIT, ABS_PRESSURE);
    if (ret != 0) {
    	printf("Error setting ABS_PRESSURE.\n");
        perror(0);
        return 1;
    }

	printf("The touchscreen event types were set successfully.\n");

    struct uinput_user_dev uidev;

    memset(&uidev, 0, sizeof(uidev));

    snprintf(uidev.name, UINPUT_MAX_NAME_SIZE, "ADS7846 Touchscreen (MOCK)");
    uidev.absmax[ABS_X] = 4095;
    uidev.absmax[ABS_Y] = 4095;
    uidev.absmax[ABS_PRESSURE] = 255;

    ret = write(ts_fd, &uidev, sizeof(uidev));

    ret = ioctl(ts_fd, UI_DEV_CREATE);

    int i = strtol("0666", 0, 8);
    if (chmod (device, i) < 0) {
        printf("Error: cannot chmod %s to 666", device);
        return 1;
    }

	return 0;
}

void ts_press(int x, int y) {

    float rx = x / 320.0;
    float ry = y / 240.0;

    int tx = 3910 - ry * 3750;
    int ty = 133 + rx * 3827;

    printf("press x:%d x:%d -> ABS_X:%d ABS_Y:%d\n", x, y, tx, ty);

    struct input_event ev[5];

    memset(ev, 0, sizeof(ev));

    ev[0].type = EV_KEY;
    ev[0].code = BTN_TOUCH;
    ev[0].value = 1;
    ev[1].type = EV_ABS;
    ev[1].code = ABS_X;
    ev[1].value = tx;
    ev[2].type = EV_ABS;
    ev[2].code = ABS_Y;
    ev[2].value = ty;
    ev[3].type = EV_ABS;
    ev[3].code = ABS_PRESSURE;
    ev[3].value = 255;
    ev[4].type = EV_SYN;
    ev[4].code = SYN_REPORT;
    ev[4].value = 0;

    write(ts_fd, ev, sizeof(ev));

}

void ts_release() {

    printf("release\n");

    struct input_event ev[2];

    memset(ev, 0, sizeof(ev));

    ev[0].type = EV_KEY;
    ev[0].code = BTN_TOUCH;
    ev[0].value = 0;
    ev[1].type = EV_SYN;
    ev[1].code = SYN_REPORT;
    ev[1].value = 0;

    write(ts_fd, ev, sizeof(ev));

}
