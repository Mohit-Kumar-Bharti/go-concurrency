package main

// TODO: Implement relaying of message with Channel Direction

func genMsg(ch1 chan<- string) {
	// send message on ch1
	ch1 <- "hello"
}

func relayMsg(ch1 <-chan string, ch2 chan<- string) {

	msg := <-ch1
	ch2 <- msg
	// recv message on ch1
	// send it on ch2
}

func main() {
	// create ch1 and ch2
	ch1 := make(chan string)
	ch2 := make(chan string)

	go genMsg(ch1)

	go relayMsg(ch1, ch2)

	v := <-ch2

	println("Received message", v)

	// spine goroutine genMsg and relayMsg

	// recv message on ch2
}
