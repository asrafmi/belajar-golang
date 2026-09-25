package context

import (
	"context"
	"fmt"
	"testing"
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
