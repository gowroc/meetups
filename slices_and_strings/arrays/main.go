package main

import (
	"fmt"
)

func main() {
	a := [3]int{1, 2, 3}
	aPtr := &a
	aPtr[0] = 3
	fmt.Println(a)
	fmt.Println(aPtr)
}
