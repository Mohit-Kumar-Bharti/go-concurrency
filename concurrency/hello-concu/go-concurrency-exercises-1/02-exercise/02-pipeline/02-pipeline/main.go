// generator() -> square() -> print

package main

import "sync"

func generator(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(cs))

	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			out <- n
		}
	}

	for _, ch := range cs {
		go output(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out

	// Implement fan-in
	// merge a list of channels to a single channel
}

func main() {
	in := generator(2, 3)

	// TODO: fan out square stage to run two instances.
	ch1 := square(in)
	ch2 := square(in)

	// TODO: fan in the results of square stages.
	merged := merge(ch1, ch2)

	// receive the values from merged channel
	// print each one, until channel is closed.
	for n := range merged {
		println(n)
	}

}
