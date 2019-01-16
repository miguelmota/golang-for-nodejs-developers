package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5}
	fmt.Println(array)

	clone := make([]int, len(array))
	copy(clone, array)
	fmt.Println(clone)

	sub := array[2:4]
	fmt.Println(sub)

	concatenated := append(array, []int{6, 7}...)
	fmt.Println(concatenated)

	prepended := append([]int{-2, -1, 0}, concatenated...)
	fmt.Println(prepended)
}
