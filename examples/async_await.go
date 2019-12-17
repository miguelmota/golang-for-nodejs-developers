package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/prometheus/common/log"
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

func main() {
	result := <-hello("bob")
	switch v := result.(type) {
	case string:
		fmt.Println(v)
	case error:
		log.Errorln(v)
	}

	result = <-hello("fail")
	switch v := result.(type) {
	case string:
		fmt.Println(v)
	case error:
		log.Errorln(v)
	}
}
