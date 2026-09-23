package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(3 * time.Second)
	timeBiasa := time.Now()

	timeBeda := <-timer.C
	fmt.Println("time biasa ", timeBiasa)
	fmt.Println("time beda ", timeBeda)
}

func TestAfter(t *testing.T) {
	timer := time.After(3 * time.Second)
	timeBiasa := time.Now()

	fmt.Println("time biasa ", timeBiasa)
	fmt.Println("time beda ", <-timer)
}

func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)
	time.AfterFunc(3*time.Second, func() {
		fmt.Println("now ->", time.Now())
		group.Done()
	})
	timeBiasa := time.Now()

	group.Wait()
	fmt.Println("time biasa ", timeBiasa)
}
