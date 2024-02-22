CC = gcc
CFLAGS = -W -Wall
TARGET = hello
OBJECTS = hello.o

all : $(TARGET)

$(TARGET) : $(OBJECTS)
	$(CC) $(CFLAGS) -o $@ $^

clean : 
	rm  *.o $(TARGET)

