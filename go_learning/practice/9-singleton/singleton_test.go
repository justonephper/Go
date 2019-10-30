package singleton

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Student struct {
}

var singletonInstance *Student
var once sync.Once

func getObj() *Student {
	once.Do(func() {
		fmt.Println("obj was created.")
		singletonInstance = new(Student)
	})
	return singletonInstance
}

func TestSingleton(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(unsafe.Pointer(getObj()))
			wg.Done()
		}()
	}
	wg.Wait()
}
