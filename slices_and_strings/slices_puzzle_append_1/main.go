package main

import "fmt"

func main() {
	s := make([]int, 0 /*length*/, 1 /*capacity*/)
	s1 := append(s, 1)
	_ = append(s, 2)
	fmt.Println(s1)
}
