CC = gcc
CFLAGS  = -g -Wall
BIN = ./bin
SRC = ./src

default: viewfb

viewfb: framebuffer.o touchscreen.o main.o
	$(CC) $(CFLAGS) -o $(BIN)/viewfb $(BIN)/viewfb.o $(BIN)/framebuffer.o $(BIN)/touchscreen.o -L/usr/X11R6/lib -lX11 -pthread

main.o: $(SRC)/main.c
	$(CC) $(CFLAGS) -c $(SRC)/main.c -o $(BIN)/main.o

framebuffer.o: $(SRC)/framebuffer.c
	$(CC) $(CFLAGS) -c $(SRC)/framebuffer.c -o $(BIN)/framebuffer.o

touchscreen.o: $(SRC)/touchscreen.c
	$(CC) $(CFLAGS) -c $(SRC)/touchscreen.c -o $(BIN)/touchscreen.o

clean:
	$(RM) $(BIN)/*
