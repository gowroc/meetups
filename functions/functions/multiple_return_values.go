package main

import (
	"fmt"
)

func multiple() (int, string, error) { // HL
	return 1, "hello", nil
}

func main() {
	a, b, c := multiple() // HL
	fmt.Println(a, b, c)
}
