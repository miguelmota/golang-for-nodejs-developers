package main

import (
	"errors"
	"fmt"
	"time"
)

func hello(name string) chan interface{} {
	ch := make(chan interface{}, 1)
	go func() {
		time.Sleep(1 * time.Second)
		if name == "fail" {
			ch <- errors.New("failed")
		} else {
			ch <- "hello " + name
		}
	}()

	return ch
}

func await(ch chan interface{}) (string, error) {
	res := <-ch
	switch v := res.(type) {
	case string:
		return v, nil
	case error:
		return "", v
	default:
		return "", errors.New("unknown")
	}
}

func main() {
	result, err := await(hello("bob"))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)

	result, err = await(hello("fail"))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}
