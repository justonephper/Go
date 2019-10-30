package util_first_response

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(i int, ch chan string) string {
	time.Sleep(time.Millisecond * 100)
	return fmt.Sprintf("%d is first response.", i)
}

func FirstResponse() {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)

	for i := 0; i < numOfRunner; i++ {
		go func(i int, ch chan string) {
			ret := runTask(i, ch)
			ch <- ret
		}(i, ch)
	}
	fmt.Println(<-ch)
}

func TestFirstResponse(t *testing.T) {
	fmt.Println("before:", runtime.NumGoroutine())
	FirstResponse()
	time.Sleep(time.Second * 1)
	fmt.Println("after:", runtime.NumGoroutine())
}
