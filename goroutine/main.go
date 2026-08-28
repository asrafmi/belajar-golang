package main

import (
	"fmt"
	"time"
)

func MainHelloWorld() {
	fmt.Println("Hello World!")
}

func main() {
	go MainHelloWorld()
	fmt.Println("lohh!")

	time.Sleep(1 * time.Second)
}
