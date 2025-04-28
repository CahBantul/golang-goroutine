package goroutine_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	var counter int
	var mutex sync.Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			mutex.Lock()
			counter++
			mutex.Unlock()
		}()
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Counter:", counter)
}

func TestRaceConditionWithChannel(t *testing.T) {
	var counter int
	var channel = make(chan struct{}, 1)

	for i := 0; i < 1000; i++ {
		go func() {
			channel <- struct{}{}
			counter++
			<-channel
		}()
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Counter:", counter)
}
