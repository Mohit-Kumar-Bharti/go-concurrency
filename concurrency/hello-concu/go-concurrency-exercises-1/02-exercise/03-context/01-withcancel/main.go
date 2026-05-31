package main

import (
	"context"
	"fmt"
)

func main() {

	// TODO: generator -  generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the goroutine once
	// they consume 5th integer value
	// so that internal goroutine
	// started by gen is not leaked.
	generator := func(ctx context.Context) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for i := 1; ; i++ {
				select {
				case out <- i:
				case <-ctx.Done():
					return
				}
			}
		}()
		return out

	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := generator(ctx)
	// Create a context that is cancellable.

	for val := range ch {
		fmt.Println(val)
		if val == 5 {
			cancel() // Cancel the context to stop the generator goroutine.
			break
		}
	}

}
