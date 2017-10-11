package combinatorics

import (
	"fmt"
	"testing"
)

func TestFactorial(t *testing.T) {
	tests := []struct {
		input  int
		output int64
		name   string
	}{
		{
			input:  -1,
			output: 1,
			name:   "returns 1 for negative numbers",
		},
		{
			input:  0,
			output: 1,
			name:   "returns 1 for 0",
		},
		{
			input:  1,
			output: 1,
			name:   "returns 1 for 1",
		},
		{
			input:  4,
			output: 24,
			name:   "returns 5 for 4",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := Factorial(test.input)
			if r != test.output {
				t.Fatalf("Expected %d, got %d", test.output, r)
			}
		})
	}
}

func BenchmarkFactorial(b *testing.B) {
	const max = 50
	for i := 0; i < max; i++ {
		b.Run(fmt.Sprintf("reult for %d", i), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				Factorial(i)
			}
		})
	}
}

func ExampleFactorial_negative() {
	fmt.Println(Factorial(-1))
	// Output: 1
}

func ExampleFactorial_zero() {
	fmt.Println(Factorial(0))
	// Output: 1
}

func ExampleFactorial_positive() {
	fmt.Println(Factorial(4))
	// Output: 24
}
