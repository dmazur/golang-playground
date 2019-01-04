package main

import "fmt"

func main() {
	var bytes int = 0

	bytes, _ = fmt.Println("Trying out Go")
	// 14 because it counts the new line character
	fmt.Println(bytes)

	bytes, _ = fmt.Print("Trying out Go")
	// 13 because it doesn't print a new line
	fmt.Print("\n")
	fmt.Println(bytes)
}
