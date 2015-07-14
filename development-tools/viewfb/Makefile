CC = gcc
CFLAGS  = -g -Wall

default: viewfb

viewfb: viewfb.o framebuffer.o touchscreen.o
	$(CC) $(CFLAGS) -o viewfb viewfb.o framebuffer.o touchscreen.o -L/usr/X11R6/lib -lX11 -pthread

viewfb.o:  viewfb.c framebuffer.h touchscreen.h
	$(CC) $(CFLAGS) -c viewfb.c

framebuffer.o:  framebuffer.c
	$(CC) $(CFLAGS) -c framebuffer.c

touchscreen.o:  touchscreen.c
	$(CC) $(CFLAGS) -c touchscreen.c

clean:
	$(RM) count *.o *~ viewfb
