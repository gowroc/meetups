package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// START OMIT
func grow(s []int) ([]int, int) {
	l := len(s)
	if cap(s) > l {
		return s[0 : l+1], l
	}
	ns := make([]int, l+1, max(2*l, 1))
	copy(ns, s)
	return ns, l
}

func intsAppend(s []int, i int) []int {
	s, oldLen := grow(s)
	s[oldLen] = i
	return s
}

func main() {
	s := []int{}
	fmt.Println("len", len(s), "cap", cap(s), SliceHeader(s), s)
	for i := 0; i < 10; i++ {
		s = intsAppend(s, i)
		fmt.Println("len", len(s), "cap", cap(s), SliceHeader(s), s)
	}
}

// END OMIT

func SliceHeader(i []int) *reflect.SliceHeader {
	return (*reflect.SliceHeader)(unsafe.Pointer(&i))
}
