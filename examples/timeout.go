package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func callback() {
	defer wg.Done()
	fmt.Println("called")
}

func main() {
	wg.Add(1)
	time.AfterFunc(1*time.Second, callback)
	wg.Wait()
}
