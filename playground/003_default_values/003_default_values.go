package main

import "fmt"

func main() {
	// "empty" values for different types
	var theString string
	printTypeAndValue(theString)

	var theInt int
	printTypeAndValue(theInt)

	var theBool bool
	printTypeAndValue(theBool)

	var theFloat32 float32
	printTypeAndValue(theFloat32)

	// type base on different type
	type notint int
	var theNotint notint
	printTypeAndValue(theNotint)
}

func printTypeAndValue(val interface{}) {
	fmt.Printf("Value: %#v Type: %T\n", val, val)
}
