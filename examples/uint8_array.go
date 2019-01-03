package main

import "fmt"

func main() {
	array := make([]byte, 10)
	fmt.Println(array)

	offset := 1

	copy(array[offset:], []byte{1, 2, 3})
	fmt.Println(array)

	sub := array[2:]
	fmt.Println(sub)

	sub2 := array[2:4]
	fmt.Println(sub2)

	fmt.Println(len(array))
}
