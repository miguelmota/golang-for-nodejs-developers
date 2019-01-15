package main

import (
	"fmt"
	"net"
)

func main() {
	ips, err := net.LookupIP("google.com")
	if err != nil {
		panic(err)
	}

	fmt.Println(ips)

	mx, err := net.LookupMX("google.com")
	if err != nil {
		panic(err)
	}

	fmt.Println(mx)

	txt, err := net.LookupTXT("google.com")
	if err != nil {
		panic(err)
	}

	fmt.Println(txt)
}
