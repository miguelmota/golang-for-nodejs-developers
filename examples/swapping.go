package main

import "fmt"

func main() {
	a := "foo"
	b := "bar"

	fmt.Println(a, b)

	b, a = a, b

	fmt.Println(a, b)
}
