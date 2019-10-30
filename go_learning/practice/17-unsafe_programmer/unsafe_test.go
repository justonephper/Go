package unsafe__test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

func TestUnsafe(t *testing.T) {
	i := 10
	f := *(*float64)(unsafe.Pointer(&i))
	t.Log(unsafe.Pointer(&i))
	t.Log(unsafe.Pointer(&f))
	t.Log(f)
}

//合理的类型转换
type MyInt int

func TestConvert(t *testing.T) {
	a := []int{1, 2, 3}
	b := *(*[]MyInt)(unsafe.Pointer(&a))

	t.Log(b)
}

//原子类型操作
func TestAtomic(t *testing.T) {
	var shareBuffer unsafe.Pointer
	writeFunc := func() {
		data := []int{}
		for i := 0; i < 100; i++ {
			data = append(data, i)
		}
		atomic.StorePointer(&shareBuffer, unsafe.Pointer(&data))
	}

	readFunc := func() {
		data := atomic.LoadPointer(&shareBuffer)
		fmt.Println(data, *(*[]int)(data))
	}
	var wg sync.WaitGroup
	writeFunc()
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				writeFunc()
				time.Sleep(time.Millisecond * 100)
			}
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				readFunc()
				time.Sleep(time.Millisecond * 100)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
