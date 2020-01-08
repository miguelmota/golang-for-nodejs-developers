package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type PersonCollection []Person

func (pc PersonCollection) Len() int {
	return len(pc)
}

func (pc PersonCollection) Swap(i, j int) {
	pc[i], pc[j] = pc[j], pc[i]
}

func (pc PersonCollection) Less(i, j int) bool {
	// asc
	return pc[i].Age < pc[j].Age
}

func main() {
	intList := []int{1, 3, 5, 9, 4, 2, 0}

	// asc
	sort.Ints(intList)
	fmt.Println(intList)
	// desc
	sort.Sort(sort.Reverse(sort.IntSlice(intList)))
	fmt.Println(intList)

	stringList := []string{"a", "d", "z", "b", "c", "y"}

	// asc
	sort.Strings(stringList)
	fmt.Println(stringList)
	// desc
	sort.Sort(sort.Reverse(sort.StringSlice(stringList)))
	fmt.Println(stringList)

	collection := []Person{
		{"Li L", 8},
		{"Json C", 3},
		{"Zack W", 15},
		{"Yi M", 2},
	}

	// asc
	sort.Sort(PersonCollection(collection))
	fmt.Println(collection)
	// desc
	sort.Sort(sort.Reverse(PersonCollection(collection)))
	fmt.Println(collection)
}
