package concurrency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func isCancelled(ch chan struct{}) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}

func cancel_1(ch chan struct{}) {
	ch <- struct{}{}
}

func cancel_all(ch chan struct{}) {
	close(ch)
}

func TestCancelByClose(t *testing.T) {
	ch := make(chan struct{})
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int, ch chan struct{}, wg *sync.WaitGroup) {
			for {
				if isCancelled(ch) {
					break
				} else {
					time.Sleep(time.Millisecond * 50)
				}
			}
			fmt.Println(i, "is cancelled!")
			wg.Done()
		}(i, ch, &wg)
	}
	cancel_all(ch)
	wg.Wait()
}
