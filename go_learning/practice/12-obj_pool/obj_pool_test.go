package obj_pool

import (
	"fmt"
	"testing"
	"time"
)

func TestObjPool(t *testing.T) {
	numOfObj := 10
	pool := NewObjPool(numOfObj)

	//尝试放置超多池子大小的对象
	//if err := pool.ReleaseObj(&ReusableObj{}); err != nil {
	//	t.Log(err)
	//}

	//获取对象
	for i := 0; i < numOfObj; i++ {
		if v, err := pool.getObj(time.Second * 1); err != nil {
			t.Error(err)
		} else {
			fmt.Printf("%T\n", v)
			if err := pool.putObj(v); err != nil {
				t.Error(err)
			}
		}
	}
}
