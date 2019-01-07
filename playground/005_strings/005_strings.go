package main

import "fmt"

func main() {
	s := "I like trains"
	fmt.Printf("Val: %v, Type: %T\n", s, s)

	// string is a slice of bytes
	bs := []byte(s)
	// byte = uint8
	fmt.Printf("Val: %v, Type: %T\n", bs, bs)

	// printting UTF-8
	for _, v := range s {
		fmt.Printf("%#U\n", v)
	}
}
