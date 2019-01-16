package main

import (
	"errors"
	"fmt"
	"runtime/debug"
)

func foo() {
	panic(errors.New("failed"))
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(string(debug.Stack()))
		}
	}()

	foo()
}
