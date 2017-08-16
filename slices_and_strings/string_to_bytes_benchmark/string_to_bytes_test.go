package slices

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"unsafe"
)

func accept(b []byte) int {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		panic(err)
	}
	return 0
}

func benchmarkNormalCast(s string) {
	accept([]byte(s))
}

func benchmarkUnsafeCast(s string) {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	slh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	accept(*(*[]byte)(unsafe.Pointer(&slh)))
}

func stringOfLength(n int) string {
	return fmt.Sprintf(`"%s"`, strings.Repeat(" ", n))
}

func BenchmarkStringConversion(b *testing.B) {
	iterations := []struct {
		name string
		test func(string)
	}{
		{"normal", benchmarkNormalCast},
		{"unsafe", benchmarkUnsafeCast},
	}

	for _, s := range []string{
		stringOfLength(10),
		stringOfLength(100),
		stringOfLength(1000),
		stringOfLength(10000),
		stringOfLength(100000),
		stringOfLength(1000000),
		stringOfLength(10000000),
	} {
		b.Run(fmt.Sprint("size=", len(s)), func(b *testing.B) {
			for _, iter := range iterations {
				b.Run(iter.name, func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						iter.test(s)
					}
				})
			}
		})
	}
}
