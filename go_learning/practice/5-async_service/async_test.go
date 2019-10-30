package __async_service

import (
	"fmt"
	"testing"
	"time"
)

func Service() string {
	time.Sleep(time.Millisecond*50)
	return "Done!"
}

func otherTask()  {
	fmt.Println("working on otherthing")
	time.Sleep(time.Millisecond*100)
	fmt.Println("task is done")
}

func TestService(t *testing.T)  {
	fmt.Println(Service())
	otherTask()
}

func asyncService() chan string{
	retchan := make(chan string,1)

	go func() {
		ret := Service()
		fmt.Println("returned result.")
		retchan<- ret
		fmt.Println("service exit")
	}()
	return retchan
}

func TestAsyncService(t *testing.T)  {
	ret := asyncService()
	otherTask()
	fmt.Println(<-ret)
	time.Sleep(time.Second*1)
}
