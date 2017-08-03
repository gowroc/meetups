package main

import "fmt"

// START OMIT
func main() {
	var ii []int
	// the first way
	ii = []int{0, 1, 2, 3}
	for i := range ii {
		if i == 0 {
			ii = ii[:2]

		}
		fmt.Println(i, ii[i])
	}
	fmt.Println("---")

	// the second way
	ii = []int{1, 2, 3}
	for i, v := range ii {
		if i == 0 {
			ii = ii[:2]

		}
		fmt.Println(i, v)
	}
}

// END OMIT
