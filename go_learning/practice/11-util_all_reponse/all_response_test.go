package util_all_reponse

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(i int) string {
	time.Sleep(time.Millisecond * 100)
	return fmt.Sprintf("%d is done\n", i)
}

func allResponse() string {
	numOfRunner := 10
	ch := make(chan string)
	for i := 0; i < numOfRunner; i++ {
		go func(i int, ch chan string) {
			ret := runTask(i)
			ch <- ret
		}(i, ch)
	}

	retStr := ""
	for j := 0; j < numOfRunner; j++ {
		retStr += <-ch
	}
	return retStr
}

func TestAllResponse(t *testing.T) {
	fmt.Println("before:", runtime.NumGoroutine())
	fmt.Println(allResponse())
	time.Sleep(time.Second * 1)
	fmt.Println("after:", runtime.NumGoroutine())
}
