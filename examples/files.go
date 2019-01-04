package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	// create file
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}

	// close file
	file.Close()

	// open file
	file, err = os.OpenFile("test.txt", os.O_RDWR, 0755)
	if err != nil {
		panic(err)
	}

	// file descriptor
	fd := file.Fd()

	// open file (using file descriptor)
	file = os.NewFile(fd, "test file")

	wbuf := []byte("hello world.")
	rbuf := make([]byte, 12)
	var off int64

	// write file
	if _, err := file.WriteAt(wbuf, off); err != nil {
		panic(err)
	}

	// read file
	if _, err := file.ReadAt(rbuf, off); err != nil {
		panic(err)
	}

	fmt.Println(string(rbuf))

	// close file (using file descriptor)
	if err := syscall.Close(int(fd)); err != nil {
		panic(err)
	}

	// delete file
	if err := os.Remove("test.txt"); err != nil {
		panic(err)
	}
}
