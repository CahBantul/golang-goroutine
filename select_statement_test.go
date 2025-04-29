package goroutine_test

import (
	"fmt"
	"testing"
	"time"
)

func TestSelectChannel(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	defer close(ch1)
	defer close(ch2)
	defer close(ch3)

	go func() {
		ch1 <- "data from channel 1"
	}()

	go func() {
		ch2 <- "data from channel 2"
	}()

	go func() {
		ch3 <- "data from channel 3"
	}()

	received := 0
	for received < 3 {
		select {
		case data := <-ch1:
			fmt.Println(data)
			received++
		case data := <-ch2:
			fmt.Println(data)
			received++
		case data := <-ch3:
			fmt.Println(data)
			received++
		}
	}
}

func TestSelectTimeout(t *testing.T) {
	ch1 := make(chan string)
	defer close(ch1)

	// go func() {
	// 	ch1 <- "data from channel 1"
	// }()

	select {
	case data := <-ch1:
		fmt.Println(data)
	case <-time.After(time.Second * 2):
		fmt.Println("timeout, No data received")
	}
}
