CC = gcc
CFLAGS  = -g -Wall
BIN = ./bin
SRC = ./src

default: viewfb

viewfb: framebuffer.o touchscreen.o viewfb.o
	$(CC) $(CFLAGS) -o $(BIN)/viewfb $(BIN)/viewfb.o $(BIN)/framebuffer.o $(BIN)/touchscreen.o -L/usr/X11R6/lib -lX11 -pthread

viewfb.o: $(SRC)/viewfb.c
	$(CC) $(CFLAGS) -c $(SRC)/viewfb.c -o $(BIN)/viewfb.o

framebuffer.o: $(SRC)/framebuffer.c
	$(CC) $(CFLAGS) -c $(SRC)/framebuffer.c -o $(BIN)/framebuffer.o

touchscreen.o: $(SRC)/touchscreen.c
	$(CC) $(CFLAGS) -c $(SRC)/touchscreen.c -o $(BIN)/touchscreen.o

clean:
	$(RM) $(BIN)/*
