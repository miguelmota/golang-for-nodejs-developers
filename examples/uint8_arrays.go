package main

import "fmt"

func main() {
	array := make([]uint8, 10)
	fmt.Println(array)

	offset := 1

	copy(array[offset:], []uint8{1, 2, 3})
	fmt.Println(array)

	sub := array[2:]
	fmt.Println(sub)

	sub2 := array[2:4]
	fmt.Println(sub2)

	fmt.Println(array)
	value := uint8(9)
	start := 5
	end := 10
	for i := start; i < end; i++ {
		array[i] = value
	}
	fmt.Println(array)

	fmt.Println(len(array))
}
