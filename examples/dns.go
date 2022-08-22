package main

import (
	"context"
	"fmt"
	"net"
	"time"
)

func main() {
	ns, err := net.LookupNS("google.com")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", ns)

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

	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10_000),
			}
			return d.DialContext(ctx, "udp", "1.1.1.1:53")
		},
	}

	ns, _ = r.LookupNS(context.Background(), "google.com")
	fmt.Printf("%s", ns)
}
