package main

import (
	"errors"
	"fmt"
)

type FooError struct {
	s string
}

func (f *FooError) Error() string {
	return f.s
}

func NewFooError(s string) error {
	return &FooError{s}
}

func main() {
	err1 := errors.New("some error")
	fmt.Println(err1)

	err2 := NewFooError("my custom error")
	fmt.Println(err2)
}
