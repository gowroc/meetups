package main

import (
	"fmt"
)

func named(p bool) (x int) {
	x = 10
	if p {
		x := 20 // HL
		return
	}
	return
}

func main() {
	fmt.Println(named(false))
	fmt.Println(named(true))
}
