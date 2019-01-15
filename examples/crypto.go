package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	hash := sha256.Sum256([]byte("hello"))

	fmt.Println(hex.EncodeToString(hash[:]))
}
