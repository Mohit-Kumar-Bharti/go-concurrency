package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// Direct call
	go fun("direct call")

	// TODO: write goroutine with different variants for function call.

	// goroutine function call

	// goroutine with anonymous function
	go func() {
		fun("goroutine with anonymous function")
	}()

	// goroutine with function value call
	f := fun
	go f("goroutine with function value call")

	// wait for goroutines to end

	// just using sleep for now
	time.Sleep(100 * time.Millisecond)

	fmt.Println("done..")
}
