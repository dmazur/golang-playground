package main

import "fmt"

type hotdog int

var x hotdog

func main() {
	fmt.Printf("The value of variable x is %v and the variable is of the type %T\n", x, x)
	x = 42
	fmt.Printf("The new value is %v\n", x)
}
