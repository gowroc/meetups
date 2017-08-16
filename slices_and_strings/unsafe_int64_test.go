package slices

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestConvertingInt64(t *testing.T) {
	var i int64 = 2
	ii := (*[2]int32)(unsafe.Pointer(&i))
	ii[1] = 1
	fmt.Println(i)
}
