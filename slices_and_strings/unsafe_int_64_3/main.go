package main

import (
	"fmt"
	"unsafe"
)

func main() {
	i3 := [3]int32{0, 1, 2}
	i9 := (*[9]int32)(unsafe.Pointer(&i3))
	fmt.Println(i3, *i9)
	i9[3] = 8
	fmt.Println(i3, *i9)
	i3[0] = 9
	fmt.Println(i3, *i9)

}
