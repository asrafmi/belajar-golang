package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Dir("work/stk/hello.go"))
	fmt.Println(path.Base("work/stk/hello.go"))
	fmt.Println(path.Ext("work/stk/hello.go"))
	fmt.Println(path.Join("work", "stk", "hello.go"))

}
