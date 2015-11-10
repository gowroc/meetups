package main

import "fmt"

func main() {
	fmt.Printf("Sum %d \n", sum(double(generate(1, 2, 3, 4, 5))))
}

// transform slice of ints into channel
func generate(nums ...int) <-chan int {
	// create channel with results
	out := make(chan int)

	// run new gouroutine
	go func() {
		for _, n := range nums {
			fmt.Printf("Generating value %d \n", n)
			out <- n
		}
		close(out)
	}()

	// return result channel immediately (it's empty at that time)
	return out
}

// read element from channel, double it and send to result channel
func double(numbers <-chan int) chan int {
	// create channel with results
	doubled := make(chan int)

	// run new gouroutine
	go func() {
		for v := range numbers {
			fmt.Printf("Doubling value %d \n", v)
			doubled <- v * 2
		}
		close(doubled)
	}()

	// return result channel immediately (it's empty at that time)
	return doubled
}

// read all elements from channel and aggregate the values
func sum(numbers chan int) int {
	acc := 0
	for v := range numbers {
		fmt.Printf("Got doubled value %d \n", v)
		acc += v
	}
	return acc
}
