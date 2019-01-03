package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
)

func writeUIntBE(buffer []byte, value, offset, byteLength int64) {
	slice := make([]byte, byteLength)
	val := new(big.Int)
	val.SetUint64(uint64(value))
	valBytes := val.Bytes()

	buf := bytes.NewBuffer(slice)
	err := binary.Write(buf, binary.BigEndian, &valBytes)
	if err != nil {
		log.Fatal(err)
	}

	slice = buf.Bytes()
	slice = slice[int64(len(slice))-byteLength : len(slice)]

	copy(buffer[offset:], slice)
}

func writeUIntLE(buffer []byte, value, offset, byteLength int64) {
	slice := make([]byte, byteLength)
	val := new(big.Int)
	val.SetUint64(uint64(value))
	valBytes := val.Bytes()

	tmp := make([]byte, len(valBytes))
	for i := range valBytes {
		tmp[i] = valBytes[len(valBytes)-1-i]
	}
	copy(slice, tmp)
	copy(buffer[offset:], slice)
}

func main() {
	buf := make([]byte, 6)
	writeUIntBE(buf, 0x1234567890ab, 0, 6)

	fmt.Println(hex.EncodeToString(buf))

	buf2 := make([]byte, 6)
	writeUIntLE(buf2, 0x1234567890ab, 0, 6)

	fmt.Println(hex.EncodeToString(buf2))
}
