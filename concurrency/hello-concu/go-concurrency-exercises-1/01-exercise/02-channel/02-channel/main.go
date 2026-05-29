package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 6; i++ {
			ch <- i
			// TODO: send iterator over channel
		}
		close(ch)
	}()

	for value := range ch {
		println("received value", value)
	}

	// TODO: range over channel to recv values

}
