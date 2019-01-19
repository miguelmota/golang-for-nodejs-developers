package main

import "fmt"

func Generator() chan string {
	c := make(chan string)

	go func() {
		c <- "hello"
		c <- "world"

		close(c)
	}()

	return c
}

func main() {
	gen := Generator()

	for true {
		value, more := <-gen
		fmt.Println(value, more)

		if !more {
			break
		}
	}

	// alternatively
	for value := range Generator() {
		fmt.Println(value)
	}
}
