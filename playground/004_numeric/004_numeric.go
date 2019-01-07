package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(1.321654 == 1.321654)

	f1 := 1.321654
	f2 := 1.321654
	fmt.Println(f1 == f2) // true
	fmt.Println(f1 != f2) // false
	fmt.Println(f1 > f2)  // false

	// int is 32/64 bits depending on architecture
	// os and arch:
	fmt.Println("OS: ", runtime.GOOS)
	fmt.Println("Arch: ", runtime.GOARCH)
}
