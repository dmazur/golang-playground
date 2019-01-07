package main

import "fmt"

func main() {
	// short declaration operator
	i0 := 0
	// var declaration operator
	var i1 = 0
	// type can be either given or read from init value
	// when given type the value will be "empty"
	var i2 int

	// these are the same
	fmt.Printf("%d (%T), %d (%T), %d (%T)\n", i0, i0, i1, i1, i2, i2)

	// weirdly, this is allowed
	var _ = 42
	var _ int
}
