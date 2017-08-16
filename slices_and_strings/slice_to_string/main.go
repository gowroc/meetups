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
	fmt.Printf("Len: %d\n", sh.Len)
	s := *((*string)(unsafe.Pointer(&sh)))
	fmt.Println(s)
}

// END OMIT
