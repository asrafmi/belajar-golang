package main

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Masuk Channel"
		fmt.Println("Selesai kirim data channel")
	}()

	data := <-channel
	fmt.Println("Isi data channel: ", data)

	time.Sleep(5 * time.Second)
}

func TestCreatePanicChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		// channel <- "Masuk Channel" // not sending the data inside channel
		fmt.Println("Selesai kirim data channel")
	}()

	data := <-channel // deadlock
	fmt.Println("Isi data channel: ", data)

	time.Sleep(5 * time.Second)
}

func TestCreateHangChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Masuk Channel"
		fmt.Println("Selesai kirim data channel")
	}()

	// data := <-channel // hang because channel not received
	fmt.Println("Isi data channel: ")

	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string, response string) {
	time.Sleep(2 * time.Second)
	channel <- response
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel, "Hello World!")

	data := <-channel
	fmt.Println(data)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Jokowi"
}

func OnlyOut(channel <-chan string) {
	time.Sleep(2 * time.Second)
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 2)
	defer close(channel)

	channel <- "Joko"
	channel <- "Widodo"

	fmt.Println(<-channel)
	fmt.Println(<-channel)

	fmt.Println("Selesai")
}

func TestBufferedChannelGoroutine(t *testing.T) {
	channel := make(chan string, 2)
	fmt.Println(cap(channel))
	fmt.Println(len(channel))
	defer close(channel)

	go func() {
		channel <- "Joko"
		fmt.Println(len(channel))
		channel <- "Widodo"
		fmt.Println(len(channel))
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Received: ", data)
	}

	fmt.Println("Done")
}
