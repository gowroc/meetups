package main // OMIT

import "fmt" // OMIT

func produce(start chan int) {
	for {
		start <- 0 // HL
	}
}

func inc(in, out chan int) {
	for {
		x := <-in // HL
		x++
		out <- x // HL
	}
}

func consume(end chan int) {
	for {
		x := <-end // HL
		fmt.Printf("%d ", x)
	}
}

func main() {
	start := make(chan int)
	end := start

	for i := 0; i < 100; i++ {
		prev := end
		end = make(chan int)
		go inc(prev, end) // HL
	}

	go produce(start) // HL
	consume(end)
}
