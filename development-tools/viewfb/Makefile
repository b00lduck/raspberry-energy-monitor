CC = gcc
CFLAGS  = -g -Wall
BIN = ./bin
SRC = ./src

default: viewfb

viewfb: framebuffer.o touchscreen.o main.o imageConverter.o exposeEventThread.o window.o
	$(CC) $(CFLAGS) -o $(BIN)/viewfb $(BIN)/main.o $(BIN)/imageConverter.o $(BIN)/exposeEventThread.o $(BIN)/framebuffer.o $(BIN)/touchscreen.o $(BIN)/window.o -L/usr/X11R6/lib -lX11 -pthread

main.o: $(SRC)/main.c
	$(CC) $(CFLAGS) -c $(SRC)/main.c -o $(BIN)/main.o

framebuffer.o: $(SRC)/framebuffer.c
	$(CC) $(CFLAGS) -c $(SRC)/framebuffer.c -o $(BIN)/framebuffer.o

touchscreen.o: $(SRC)/touchscreen.c
	$(CC) $(CFLAGS) -c $(SRC)/touchscreen.c -o $(BIN)/touchscreen.o

imageConverter.o: $(SRC)/imageConverter.c
	$(CC) $(CFLAGS) -c $(SRC)/imageConverter.c -o $(BIN)/imageConverter.o

exposeEventThread.o: $(SRC)/exposeEventThread.c
	$(CC) $(CFLAGS) -c $(SRC)/exposeEventThread.c -o $(BIN)/exposeEventThread.o

window.o: $(SRC)/window.c
	$(CC) $(CFLAGS) -c $(SRC)/window.c -o $(BIN)/window.o

clean:
	$(RM) $(BIN)/*
