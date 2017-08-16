package main

import (
	"fmt"
)

func main() {
	var ii [3]int

	// the first way
	ii = [3]int{1, 2, 3}
	for i := range ii {
		ii[1] = 4
		fmt.Println(i, ii[i])
	}
	fmt.Println("---")

	// the second way
	ii = [3]int{1, 2, 3}
	for i, v := range ii {
		ii[1] = 4
		fmt.Println(i, ii[i], v)
	}
}
