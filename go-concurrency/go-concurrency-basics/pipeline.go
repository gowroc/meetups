package main

import "fmt"

func main() {
	fmt.Printf("Sum %d \n", sum(double(generate(1, 2, 3, 4, 5))))
}

func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			fmt.Printf("Generating value %d \n", n)
			out <- n
		}
		close(out)
	}()
	return out
}

func double(numbers <-chan int) chan int {
	doubled := make(chan int)
	go func() {
		for v := range numbers {
			fmt.Printf("Doubling value %d \n", v)
			doubled <- v * 2
		}
		close(doubled)
	}()
	return doubled
}

func sum(numbers chan int) int {
	acc := 0
	for v := range numbers {
		fmt.Printf("Got doubled value %d \n", v)
		acc += v
	}
	return acc
}
