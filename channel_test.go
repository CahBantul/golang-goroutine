package goroutine_test

import (
	"fmt"
	"testing"
)

func TestChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		channel <- "Hello world goroutine"
		fmt.Println("Selesai mengirim data ke channel")
	}()

	data := <-channel
	fmt.Println("Received:", data)
}
