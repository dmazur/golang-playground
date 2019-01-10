package main

import "fmt"

func main() {
	a := 42
	b := 1337
	g := a == b
	h := a <= b
	i := a >= b
	j := a != b
	k := a < b
	l := a > b

	fmt.Printf("a = %d, b = %d\n", a, b)
	fmt.Println("a == b ?", g)
	fmt.Println("a <= b ?", h)
	fmt.Println("a >= b ?", i)
	fmt.Println("a != b ?", j)
	fmt.Println("a < b ?", k)
	fmt.Println("a > b ?", l)
}
