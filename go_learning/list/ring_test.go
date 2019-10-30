package list

import (
	"container/ring"
	"fmt"
	"testing"
)

func TestRing(t *testing.T)  {
	r := ring.New(10)

	for i:= 0 ; i < r.Len() ; i++ {
		r.Value = i
		r = r.Next()
	}

	for i:= 0; i < r.Len() ; i++ {
		fmt.Print(r.Value)
		r = r.Next()
	}
}
