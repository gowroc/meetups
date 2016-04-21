package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args) // HL
	os.Exit(42)          // HL
}
