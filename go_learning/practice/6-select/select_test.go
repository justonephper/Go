package __select

import (
	"fmt"
	"testing"
	"time"
)

func Service() string {
	time.Sleep(time.Millisecond*50)
	return "Done!"
}

func AsyncService() chan string{
	retCh := make(chan string,1)
	go func() {
		ret := Service()
		fmt.Println("return result.")
		retCh<-ret
		fmt.Println("service exit")
	}()
	return retCh
}

func TestSelect(t *testing.T)  {
	select {
	case ret := <-AsyncService():
		t.Log(ret)
		case <-time.After(time.Millisecond*100):
			t.Error("time out")
	}
}
