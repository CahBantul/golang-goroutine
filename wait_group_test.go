package goroutine_test

import (
	"fmt"
	"sync"
	"testing"
)

func runTask(number int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Runing task", number)
}

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go runTask(i, &wg)
	}

	wg.Wait()
	fmt.Println("All task is completed")
}
