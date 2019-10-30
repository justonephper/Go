package concurrency

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func Cancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancelByContext(t *testing.T) {
	var ctx, cancel = context.WithCancel(context.Background())
	var wg sync.WaitGroup
	numOfJob := 10

	for i := 0; i < numOfJob; i++ {
		wg.Add(1)
		go func(i int, ctx context.Context) {
			for {
				if Cancelled(ctx) {
					break
				} else {
					time.Sleep(time.Millisecond * 50)
				}
			}
			fmt.Println(i, " is cancelled")
			wg.Done()
		}(i, ctx)
	}
	cancel()
	wg.Wait()
}
