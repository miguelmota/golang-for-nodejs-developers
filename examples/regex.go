package main

import (
	"fmt"
	"regexp"
)

func main() {
	input := "foobar"
	re := regexp.MustCompile(`(?i)foo(.*)`)
	replaced := re.ReplaceAllString(input, "qux$1")
	fmt.Println(replaced)

	re = regexp.MustCompile(`(?i)o{2}`)
	match := re.Match([]byte(input))
	fmt.Println(match)

	input = "111-222-333"
	re = regexp.MustCompile(`(?i)([0-9]+)`)
	matches := re.FindAllString(input, -1)
	fmt.Println(matches)
}
