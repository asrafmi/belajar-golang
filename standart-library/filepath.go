package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Dir("work/stk/hello.go"))
	fmt.Println(filepath.Base("work/stk/hello.go"))
	fmt.Println(filepath.Ext("work/stk/hello.go"))
	fmt.Println(filepath.IsAbs("/Users/asrafff/Work/porto/belajar-golang/standart-library"))
	fmt.Println(filepath.IsLocal("work/stk/hello.go"))
	fmt.Println(filepath.Join("work", "stk", "hello.go"))
}
