package main

import "fmt"

func main() {
	array := []byte{1, 2}

	if array != nil {
		fmt.Println("array exists")
	}

	if len(array) == 2 {
		fmt.Println("length is 2")
	} else if len(array) == 1 {
		fmt.Println("length is 1")
	} else {
		fmt.Println("length is other")
	}

	// closest thing to ternary operator
	isOddLength := "no"
	if len(array)%2 == 1 {
		isOddLength = "yes"
	}

	fmt.Println(isOddLength)
}
