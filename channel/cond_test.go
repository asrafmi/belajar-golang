package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()

	cond.L.Lock()
	cond.Wait()

	fmt.Println("Here ->", value)

	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			group.Add(1)
			go WaitCondition(i)
		}
	}

	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		if i%2 == 0 {
	// 			time.Sleep(1 * time.Second)
	// 			cond.Signal()
	// 		}

	// 		continue
	// 	}
	// }()

	go func() {
		time.Sleep(1 * time.Second)
		cond.Broadcast()
	}()

	group.Wait()
}
