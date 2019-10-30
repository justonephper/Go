package channel_close

import (
	"fmt"
	"sync"
	"testing"
)

func dataProductor(ch chan int,wg *sync.WaitGroup) {
	go func() {
		for i:=0;i<10;i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()
}

func dataReceive(ch chan int,wg *sync.WaitGroup)  {
	go func() {
		for {
			if val,ok := <-ch;ok{
				fmt.Println(val)
			} else {
				break
			}
		}
		wg.Done()
	}()
}

func TestChanClose(t *testing.T)  {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProductor(ch,&wg)

	wg.Add(1)
	dataReceive(ch,&wg)

	wg.Wait()
}