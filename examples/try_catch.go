package main

import (
	"errors"
	"fmt"
)

func foo(fail bool) error {
	if fail {
		return errors.New("my error")
	}

	return nil
}

func main() {
	err := foo(true)
	if err != nil {
		fmt.Printf("caught error: %s\n", err.Error())
	}
}
