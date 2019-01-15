package main

import "fmt"

type Obj struct {
	SomeProperties map[string]string
}

func NewObj() *Obj {
	return &Obj{
		SomeProperties: map[string]string{
			"foo": "bar",
		},
	}
}

func (o *Obj) SomeMethod(prop string) string {
	return o.SomeProperties[prop]
}

func main() {
	obj := NewObj()

	item := obj.SomeProperties["foo"]
	fmt.Println(item)

	item = obj.SomeMethod("foo")
	fmt.Println(item)
}
