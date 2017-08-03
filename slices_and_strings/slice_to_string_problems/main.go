package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// START OMIT
func main() {
	bs := []byte("hello")
	bsh := (*reflect.SliceHeader)(unsafe.Pointer(&bs))
	sh := reflect.StringHeader{
		Data: bsh.Data,
		Len:  bsh.Len,
	}
	s := *((*string)(unsafe.Pointer(&sh)))
	fmt.Println(s)
	bs[1] = 'a'
	fmt.Println(s)
}

// END OMIT
