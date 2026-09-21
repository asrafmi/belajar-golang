package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsync(group *sync.WaitGroup, i int) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hii ke", i)
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 1; i <= 100; i++ {
		go RunAsync(group, i)
	}

	group.Wait()
	fmt.Println("Complete")
	// time.Sleep(3 * time.Second) // not needed anymore
}
