package main

func main() {
	// explicit
	var foo string = "foo"

	// type inferred
	var bar = "foo"

	// shorthand
	baz := "bar"

	// constant
	const qux = "qux"

	_ = foo
	_ = bar
	_ = baz
	_ = qux
}
