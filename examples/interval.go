package main

import (
	"fmt"
	"time"
)

func callback(i int) {
	fmt.Println("called", i)
}

func main() {
	ticker := time.NewTicker(1 * time.Second)

	i := 0
	for range ticker.C {
		callback(i)

		if i == 3 {
			ticker.Stop()
			break
		}

		i++
	}
}
