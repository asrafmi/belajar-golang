package context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println("background", background)

	todo := context.TODO()
	fmt.Println("todo", todo)

	contextB := context.WithValue(background, "b", "B")
	fmt.Println(contextB)
	contextC := context.WithValue(background, "c", "C")
	fmt.Println(contextC)

	contextD := context.WithValue(contextB, "d", "D")
	fmt.Println(contextD)
	contextE := context.WithValue(contextB, "e", "E")
	fmt.Println(contextE)

	contextF := context.WithValue(contextE, "f", "F")
	fmt.Println(contextF)
	contextG := context.WithValue(contextF, "g", "G")
	fmt.Println(contextG)

	fmt.Println(contextG.Value("g"))
	fmt.Println(contextG.Value("f"))
	fmt.Println(contextG.Value("d"))   // nil
	fmt.Println(background.Value("d")) // nil
}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutine Before", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)
	fmt.Println("Total Goroutine mid", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter", n)

		if n == 10 {
			break
		}
	}
	cancel()
	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine After", runtime.NumGoroutine())
}

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Total Goroutine Before", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := CreateCounter(ctx)
	fmt.Println("Total Goroutine mid", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter", n)

		if n == 10 {
			break
		}
	}
	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine After", runtime.NumGoroutine())
}

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Total Goroutine Before", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(3*time.Second))
	defer cancel()

	destination := CreateCounter(ctx)
	fmt.Println("Total Goroutine mid", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter", n)

		if n == 10 {
			break
		}
	}
	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine After", runtime.NumGoroutine())
}
