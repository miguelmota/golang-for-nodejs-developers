package main

import "fmt"

type Obj struct {
	SomeProperty string
}

func NewObj(someProperty string) *Obj {
	return &Obj{
		SomeProperty: someProperty,
	}
}

func (o *Obj) SomeMethod(prop string) string {
	if prop == "SomeProperty" {
		return o.SomeProperty
	}

	return ""
}

func main() {
	obj := NewObj("bar")

	item := obj.SomeProperty
	fmt.Println(item)

	item = obj.SomeMethod("SomeProperty")
	fmt.Println(item)
}
