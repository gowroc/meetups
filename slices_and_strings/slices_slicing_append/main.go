package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// START OMIT

func tryCopyInplace(from, to []int) []int {
	return append(to[0:0], from...)
}

func main() {
	s := []int{0, 1, 2}
	s1 := tryCopyInplace([]int{4, 5, 6}, s)
	fmt.Println(s, s1, SliceHeader(s), SliceHeader(s1))
	s2 := tryCopyInplace([]int{8, 9, 10, 11}, s)
	fmt.Println(s, s2, SliceHeader(s), SliceHeader(s2))
}

// END OMIT

func SliceHeader(i []int) *reflect.SliceHeader {
	return (*reflect.SliceHeader)(unsafe.Pointer(&i))
}
