package main

func main() {
	// primitives
	var myBool bool = true
	var myInt int = 10
	var myInt8 int8 = 10
	var myInt16 int16 = 10
	var myInt32 int32 = 10
	var myInt64 int64 = 10
	var myUint uint = 10
	var myUint8 uint8 = 10
	var myUint16 uint16 = 10
	var myUint32 uint32 = 10
	var myUint64 uint64 = 10
	var myUintptr uintptr = 10
	var myFloat32 float32 = 10.5
	var myFloat64 float64 = 10.5
	var myComplex64 complex64 = -1 + 10i
	var myComplex128 complex128 = -1 + 10i
	var myString string = "foo"
	var myByte byte = 10  // alias to uint8
	var myRune rune = 'a' // alias to int32

	// composite types
	var myStruct struct{} = struct{}{}
	var myArray []string = []string{}
	var myMap map[string]int = map[string]int{}
	var myFunction func() = func() {}
	var myChannel chan bool = make(chan bool)
	var myInterface interface{} = nil
	var myPointer *int = new(int)

	_ = myBool
	_ = myInt
	_ = myInt8
	_ = myInt16
	_ = myInt32
	_ = myInt64
	_ = myUint
	_ = myUint8
	_ = myUint16
	_ = myUint32
	_ = myUint64
	_ = myUintptr
	_ = myFloat32
	_ = myFloat64
	_ = myComplex64
	_ = myComplex128
	_ = myString
	_ = myByte
	_ = myRune
	_ = myStruct
	_ = myArray
	_ = myMap
	_ = myFunction
	_ = myChannel
	_ = myInterface
	_ = myPointer
}
