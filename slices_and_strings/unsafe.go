package slices

import (
	"reflect"
	"unsafe"
)

// UnsafeByteArrayToStr uses unsafe to convert byte array into string. Supplied array cannot be
// altered after this functions is called
func UnsafeByteArrayToStr(b []byte) string {
	if b == nil {
		return ""
	}
	return *(*string)(unsafe.Pointer(&b))
}

// UnsafeStrToByteArray uses unsafe to convert string into byte array. Returned array cannot be
// altered after this functions is called
func UnsafeStrToByteArray(s string) []byte {
	strHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))

	var b []byte
	byteHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	byteHeader.Data = strHeader.Data

	// need to take the length of s here to ensure s is live until after we update b's Data
	// field since the garbage collector can collect a variable once it is no longer used
	// not when it goes out of scope, for more details see https://github.com/golang/go/issues/9046
	l := len(s)
	byteHeader.Len = l
	byteHeader.Cap = l
	return b
}
