package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
)

func main() {
	data := []byte("hello world\n")

	var compressed bytes.Buffer
	w := gzip.NewWriter(&compressed)
	w.Write(data)
	w.Close()

	fmt.Println(compressed.Bytes())

	var decompressed bytes.Buffer
	r, err := gzip.NewReader(&compressed)
	if err != nil {
		panic(err)
	}

	_, err = decompressed.ReadFrom(r)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(decompressed.Bytes()))
}
