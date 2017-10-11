package combinatorics

// Fibonacci computes n-th element of Fibonacci sequence.
func Fibonacci(n int) int64 {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	return fibHelper(n-1, 1, 1)
}

func fibHelper(n int, prev int64, this int64) int64 {
	if n == 0 {
		return this
	}
	return fibHelper(n-1, this, this+prev)
}
