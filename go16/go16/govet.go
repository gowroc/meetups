package main

import "fmt"

func main() {
	f := func() int {
		return 1
	}

	fmt.Printf("Value should be here: %d", f)
}
