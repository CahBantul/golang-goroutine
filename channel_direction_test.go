package goroutine_test

import (
	"fmt"
	"testing"
)

func sendOnly(ch chan<- string) {
	fmt.Println("send data")
	ch <- "Data dikirim"
}

func receiveOnly(ch <-chan string) {
	data := <-ch
	fmt.Println("received:", data)
}

func TestChannelDirection(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go sendOnly(channel)
	receiveOnly(channel)
}
