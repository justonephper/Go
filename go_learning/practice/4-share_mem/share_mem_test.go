package __share_mem

import (
	"sync"
	"testing"
	"time"
)

//强制等待，弊端是，不知道程序什么时间结束
func TestCounter(t *testing.T)  {
	counter := 0
	for i:=0;i<5000 ;i++  {
		go func() {
			counter++
		}()
	}

	time.Sleep(time.Millisecond*50)
	t.Log("counter:",counter)
}

//线程安全的计算
func TestCounterThreadSafe(t *testing.T)  {
	var mut sync.Mutex
	counter := 0
	for i:=0;i<5000;i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(time.Millisecond*50)
	t.Log("counter:",counter)
}

//程序执行结束，自动回收
func TestCounterWaiteGroup(t *testing.T)  {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0

	for i:=0;i<5000 ;i++  {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Log("counter:",counter)
}