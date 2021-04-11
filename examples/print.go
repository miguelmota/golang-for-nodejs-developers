package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("print to stdout")
	fmt.Printf("format %s %v\n", "example", 1)
	fmt.Fprintf(os.Stderr, "print to stderr")
}
