package main

import (
	"fmt"
	"time"
)

func main() {
	nowUnix := time.Now().Unix()
	fmt.Println(nowUnix)

	datestr := "2019-01-17T09:24:23+00:00"
	date, err := time.Parse("2006-01-02T15:04:05Z07:00", datestr)
	if err != nil {
		panic(err)
	}

	fmt.Println(date.Unix())
	fmt.Println(date.String())

	futureDate := date.AddDate(0, 0, 14)
	fmt.Println(futureDate.String())

	formatted := date.Format("01/02/2006")
	fmt.Println(formatted)
}
