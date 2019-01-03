package main

import (
	"encoding/json"
	"fmt"
)

type T struct {
	Foo string `json:"foo"`
}

func main() {
	jsonstr := `{"foo":"bar"}`

	t := new(T)
	err := json.Unmarshal([]byte(jsonstr), t)
	if err != nil {
		panic(err)
	}

	fmt.Println(t)

	marshalled, err := json.Marshal(t)
	jsonstr = string(marshalled)
	fmt.Println(jsonstr)
}
