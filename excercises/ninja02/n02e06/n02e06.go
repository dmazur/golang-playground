package main

import "fmt"

const (
	t0      = 2019 - iota
	tminus1 = 2019 - iota
	tminus2 = 2019 - iota
	tminus3 = 2019 - iota
)

func main() {
	fmt.Println("t0 =", t0)
	fmt.Println("tminus1 =", tminus1)
	fmt.Println("tminus2 =", tminus2)
	fmt.Println("tminus3 =", tminus3)
}
