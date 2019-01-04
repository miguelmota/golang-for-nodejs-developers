package main

import (
	"fmt"

	"github.com/go-shadow/moment"
)

func main() {
	now := moment.New().Now().Unix()
	fmt.Println(now)
}
