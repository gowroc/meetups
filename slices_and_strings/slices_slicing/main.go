package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := make([]int, 0, 20)
	fmt.Println("len", len(s), "cap", cap(s), SliceHeader(s))
	s = s[0:10]
	fmt.Println("len", len(s), "cap", cap(s), SliceHeader(s))
	s = s[1:10]
	fmt.Println("len", len(s), "cap", cap(s), SliceHeader(s))
}

func SliceHeader(i []int) *reflect.SliceHeader {
	return (*reflect.SliceHeader)(unsafe.Pointer(&i))
}
