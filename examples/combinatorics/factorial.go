package combinatorics

// Factorial is defined as a number of permutations on n elements.
func Factorial(n int) int64 {
	if n <= 0 {
		return 1
	}
	return factorialHelper(n, 1)
}

func factorialHelper(n int, acc int64) int64 {
	if n == 1 {
		return acc
	}
	return factorialHelper(n-1, acc*int64(n))
}
