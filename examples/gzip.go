package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
)

func main() {
	data := []byte("hello world")

	compressed := new(bytes.Buffer)
	w := gzip.NewWriter(compressed)
	if _, err := w.Write(data); err != nil {
		panic(err)
	}
	if err := w.Close(); err != nil {
		panic(err)
	}

	fmt.Println(compressed.Bytes())

	decompressed := new(bytes.Buffer)
	r, err := gzip.NewReader(compressed)
	if err != nil {
		panic(err)
	}

	_, err = decompressed.ReadFrom(r)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(decompressed.Bytes()))
}
