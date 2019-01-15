package main

import "errors"

func foo() {
	panic(errors.New("failed"))
}

func main() {
	foo()
}
