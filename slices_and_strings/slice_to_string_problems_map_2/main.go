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
	m := map[string]int{
		s: 1,
	}
	for i := 0; i < 10; i++ {
		m[fmt.Sprint(i)] = i
	}
	bs[1] = 'a'
	_, hasHello := m["hello"]
	_, hasHallo := m["hallo"]
	fmt.Printf("hello: %v, hallo: %v", hasHello, hasHallo)
}

// END OMIT
