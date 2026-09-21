package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	group := sync.WaitGroup{}
	pool := sync.Pool{
		New: func() interface{} {
			return "Default"
		},
	}
	students := []string{
		"Jokowi",
		"Prabowo",
		"Gibran",
	}

	for _, i := range students {
		pool.Put(i)
	}

	for i := 1; i <= 100; i++ {
		go func() {
			data := pool.Get()
			fmt.Println("Student", data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	group.Wait()
	fmt.Println("Finished")
}
