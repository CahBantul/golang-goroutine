package goroutine_test

import (
	"fmt"
	"testing"
)

func sayHello() {
	fmt.Println("Hello world goroutine")
}

func TestGoroutine(t *testing.T) {
	go sayHello()
	fmt.Println("Hello world from unit test")
}

func displayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 1000; i++ {
		go displayNumber(i)
	}
}
