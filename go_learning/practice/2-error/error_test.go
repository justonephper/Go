package error_test

import (
	"fmt"
	"testing"
)

func deferFn()  {
	//fmt.Println("panic happen")
	if err:=recover();err!=nil{
		fmt.Println("panic recover from ",err)
	}
}

func TestPanicVxExit(t *testing.T)  {
	defer deferFn()

	fmt.Println("start")
	panic("something is wrong")
}
