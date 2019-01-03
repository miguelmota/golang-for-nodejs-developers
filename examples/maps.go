package main

import "fmt"

func main() {
	myMap := make(map[string]string)

	myMap["foo"] = "bar"

	item, found := myMap["foo"]
	fmt.Println(found)
	fmt.Println(item)

	delete(myMap, "foo")

	item, found = myMap["foo"]
	fmt.Println(found)
	fmt.Println(item)
}
