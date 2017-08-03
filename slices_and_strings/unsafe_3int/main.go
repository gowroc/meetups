package main

import (
	"fmt"
	"unsafe"
)

type int3 struct {
	int1, int2, int3 int32
}

func main() {
	a := [3]int32{1, 2, 3}
	fmt.Println(a)
	i := (*int3)(unsafe.Pointer(&a))
	fmt.Println(*i)
	i.int2 = 4
	fmt.Println(a)
}
