package main

import (
	"fmt"
	"testing"
	"time"
)

func HelloWorld() {
	fmt.Println("Hello World!")
}

func TestHelloWorld(t *testing.T) {
	go HelloWorld()
	// HelloWorld()
	fmt.Println("lohh!")

	// time.Sleep(1 * time.Second)
}

func PrintNumber(number int) {
	fmt.Println("Number => ", number)
}
func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go PrintNumber(i)
	}

	time.Sleep(5 * time.Second)
}
