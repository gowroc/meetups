package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i64 int64 = 2 << 31
	fmt.Println(i64)
	i32 := (*int32)(unsafe.Pointer(&i64))
	*i32 = 1
	fmt.Println(i64)
}
