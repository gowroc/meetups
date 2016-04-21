package main

import "fmt"

func main() {
	text := "external variable" // HL

	func() {
		fmt.Println(text) // HL
	}()

	func(arg string) {
		fmt.Println(arg) // HL
	}(text)
}
