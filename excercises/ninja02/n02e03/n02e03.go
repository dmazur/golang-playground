package main

import "fmt"

const (
	answer        = 42
	intanswer int = 42
)

func main() {
	const (
		canidothis = 1337
	)

	fmt.Println("answer is", answer)
	fmt.Println("intanswer is", intanswer)
	fmt.Println("canidothis is", canidothis)
}
