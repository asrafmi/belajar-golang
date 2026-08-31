package main

import (
	"fmt"
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
