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
	res := make([]string, len(ch))
	for i, c := range ch {
		wg.Add(1)
		go func(j int, s chan string) {
			res[j] = <-s
			wg.Done()
		}(i, c)
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
