package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i int64 = 1 + 2*2<<31
	fmt.Println(i)
	ii := (*[2]int32)(unsafe.Pointer(&i))
	fmt.Println(*ii)
	ii[0] = 3
	fmt.Println(i)
}
