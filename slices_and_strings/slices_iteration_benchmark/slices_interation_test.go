package slices

import (
	"fmt"
	"testing"
)

type big [10]int64

var bigPtrs []*big
var bigs []big

func init() {
	bigs = make([]big, 100000000)
	bigPtrs = make([]*big, 100000000)
	for i := range bigPtrs {
		bigPtrs[i] = &big{}
	}
}

func accept(*big) int {
	return 0
}

func benchmarkNormalSlice(n int) {
	for i, b := range bigs {
		if i > n {
			return
		}
		accept(&b)
	}
}

func benchmarkNormalSlice2(n int) {
	for i := range bigs {
		if i > n {
			return
		}
		accept(&bigs[i])
	}
}

func benchmarkPtrSlice(n int) {
	for i, b := range bigPtrs {
		if i > n {
			return
		}
		accept(b)
	}
}

func BenchmarkSliceIteration(b *testing.B) {
	iterations := []struct {
		name string
		test func(int)
	}{
		{"normal", benchmarkNormalSlice},
		{"normal2", benchmarkNormalSlice2},
		{"pointers", benchmarkPtrSlice},
	}

	for _, n := range []int{
		10,
		100,
		10000,
		100000,
	} {
		b.Run(fmt.Sprint("size=", n), func(b *testing.B) {
			for _, iter := range iterations {
				b.Run(iter.name, func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						iter.test(n)
					}
				})
			}
		})
	}
}
