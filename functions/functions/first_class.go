package main

import "fmt"

type OPER func(int, int) int // HL

func sum(x, y int) int {
	return x + y
}

func mul(x, y int) int {
	return x * y
}

func main() {
	var op OPER // HL
	op = sum // HL
	fmt.Println(op(2, 3))
	op = mul // HL
	fmt.Println(op(2, 3))
}
