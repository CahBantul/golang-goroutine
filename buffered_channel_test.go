package goroutine_test

import (
	"fmt"
	"sync"
	"testing"
)

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "Fardan"
	channel <- "Nozami"
	channel <- "Ajitama"

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}

func TestBufferedChannelWithWaitGroup(t *testing.T) {
	channel := make(chan int, 5)
	defer close(channel)

	var wg sync.WaitGroup
	wg.Add(2)

	// producer
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Println("Sending data to channel:", i)
			channel <- i
		}
	}()

	// consumer
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			data := <-channel
			fmt.Println("Receiving data from channel:", data)
		}
	}()

	wg.Wait()
	fmt.Println("Done")
}
