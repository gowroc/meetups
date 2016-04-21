package main

import (
	"fmt"
)

func param(param string) string {
	fmt.Println("evaluation of " + param)
	return param
}

func main() {
	defer fmt.Println(param("one")) // HL
	defer fmt.Println(param("two")) // HL
	fmt.Println("main")
}
