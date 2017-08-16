package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := []int{}
	fmt.Println("len", len(s), "cap", cap(s), SliceHeader(s))
	for i := 0; i < 10; i++ {
		s = append(s, i)
		fmt.Println("len", len(s), "cap", cap(s), SliceHeader(s))
	}
}

func SliceHeader(i []int) *reflect.SliceHeader {
	return (*reflect.SliceHeader)(unsafe.Pointer(&i))
}
