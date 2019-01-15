package main

import "fmt"

func main() {
	map1 := make(map[string]string)

	map1["foo"] = "bar"

	item, found := map1["foo"]
	fmt.Println(found)
	fmt.Println(item)

	delete(map1, "foo")

	item, found = map1["foo"]
	fmt.Println(found)
	fmt.Println(item)

	map2 := make(map[string]int)
	map2["foo"] = 100
	map2["bar"] = 200
	map2["baz"] = 300

	for key, value := range map2 {
		fmt.Println(key, value)
	}
}
