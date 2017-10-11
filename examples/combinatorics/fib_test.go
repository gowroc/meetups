package combinatorics

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		input  int
		output int64
		name   string
	}{
		{
			input:  -1,
			output: 0,
			name:   "returns 0 for negative numbers",
		},
		{
			input:  0,
			output: 1,
			name:   "returns 0 for 0",
		},
		{
			input:  1,
			output: 1,
			name:   "returns 1 for 1",
		},
		{
			input:  4,
			output: 5,
			name:   "returns 5 for 4",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := Fibonacci(test.input)
			if r != test.output {
				t.Fatalf("Expected %d, got %d", test.output, r)
			}
		})
	}
}

func BenchmarkFibonacci(b *testing.B) {
	const max = 50
	for i := 0; i < max; i++ {
		b.Run(fmt.Sprintf("reult for %d", i), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				Fibonacci(i)
			}
		})
	}
}

func ExampleFibonacci_negative() {
	fmt.Println(Fibonacci(-1))
	// Output: 1
}

func ExampleFibonacci_zero() {
	fmt.Println(Fibonacci(0))
	// Output: 1
}

func ExampleFibonacci_positive() {
	fmt.Println(Fibonacci(4))
	// Output: 5
}
