package main

import (
	"fmt"
)

type MyEmitter map[string]chan string

func main() {
	myEmitter := MyEmitter{}
	myEmitter["my-event"] = make(chan string)
	myEmitter["my-other-event"] = make(chan string)

	go func() {
		for {
			select {
			case msg := <-myEmitter["my-event"]:
				fmt.Println(msg)
			case msg := <-myEmitter["my-other-event"]:
				fmt.Println(msg)
			}
		}
	}()

	myEmitter["my-event"] <- "hello world"
	myEmitter["my-other-event"] <- "hello other world"
}
