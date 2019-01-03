package main

import "fmt"

func main() {
	array := []byte{1, 2, 3, 4, 5}
	fmt.Println(array)

	clone := make([]byte, len(array))
	copy(clone, array)
	fmt.Println(clone)

	sub := array[2:4]
	fmt.Println(sub)

	concatenated := append(array, []byte{6, 7}...)
	fmt.Println(concatenated)
}
