package main

func main() {
	ch := make(chan int)
	go func(a, b int) {
		c := a + b
		ch <- c
	}(1, 2)

	res := <-ch
	println("computed value", res)
	// TODO: get the value computed from goroutine
	// fmt.Printf("computed value %v\n", c)
}
