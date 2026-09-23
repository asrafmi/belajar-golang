package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGoMaxProcs(t *testing.T) {
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}
	cpu := runtime.NumCPU()
	fmt.Println("cpu", cpu)

	runtime.GOMAXPROCS(20)
	thread := runtime.GOMAXPROCS(-1)
	fmt.Println("thread", thread)

	goroutine := runtime.NumGoroutine()
	fmt.Println("goroutine", goroutine)

}
