package main

import (
	"fmt"
	"sync"
	"time"
)

func myPromise(value string) chan string {
	ch := make(chan string, 1)
	go func() {
		time.Sleep(1 * time.Second)
		ch <- "resolved: " + value
	}()

	return ch
}

func promiseAll(ch ...chan string) []string {
	var wg sync.WaitGroup
	var res []string
	for _, c := range ch {
		wg.Add(1)
		res = append(res, <-c)
		wg.Done()
	}

	wg.Wait()
	return res
}

func main() {
	res := <-myPromise("foo")
	fmt.Println(res)

	all := promiseAll(
		myPromise("A"),
		myPromise("B"),
		myPromise("C"),
	)

	fmt.Println(all)
}
