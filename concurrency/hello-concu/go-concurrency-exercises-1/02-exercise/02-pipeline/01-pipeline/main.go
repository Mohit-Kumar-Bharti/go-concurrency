package main

// TODO: Build a Pipeline
// generator() -> square() -> print

// generator - convertes a list of integers to a channel
func generator(nums ...int) <-chan int {
	ch := make(chan int)

	go func() {
		for _, n := range nums {
			ch <- n
		}
		close(ch)
	}()

	return ch

}

// square - receive on inbound channel
// square the number
// output on outbound channel
func square(ch <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range ch {
			out <- n * n
		}
		close(out)
	}()

	return out

}

func main() {
	// set up the pipeline

	// run the last stage of pipeline
	// receive the values from square stage
	// print each one, until channel is closed.

	for n := range square(generator(2, 3, 4)) {
		println(n)
	}

}
