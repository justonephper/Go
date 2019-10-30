package obj_pool

import (
	"errors"
	"time"
)

type ReusableObj struct {
}

type ObjPool struct {
	bufChan chan *ReusableObj
}

func NewObjPool(numOfObj int) *ObjPool {
	pool := new(ObjPool)
	pool.bufChan = make(chan *ReusableObj, numOfObj)

	for i := 0; i < numOfObj; i++ {
		pool.bufChan <- &ReusableObj{}
	}
	return pool
}

//获取对象
func (p *ObjPool) getObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeout):
		return nil, errors.New("time out")
	}
}

//放回对象
func (p *ObjPool) putObj(obj *ReusableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}
