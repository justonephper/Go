package list

import (
	"container/list"
	"fmt"
	"testing"
)

func TestMove(t *testing.T) {
  l1 := list.New()
  l2 := list.New()

  l1.PushBack("name1")
  l2.PushBack("age1")

  //l1.PushBackList(l2)
  l1.PushFrontList(l2)

  for e:=l1.Front();e != nil; e=e.Next() {
    fmt.Println(e.Value)
  }

}

//func TestListInsert(t *testing.T) {
//	l := list.New()
//	l.PushBack("name1")
//	fmt.Println(l.Back().Value)
//
//	l1 := &list.List{}
//	l1.PushBack("pwd1")
//	fmt.Println(l1.Back().Value)
//
//	l2 := new(list.List)
//	l2.PushBack("age1")
//	fmt.Println(l2.Back().Value)
//}

//func TestList(t *testing.T) {
//    //声明
//	list := new(list.List)
//
//	//循环赋值
//	for i := 0; i < 10; i++ {
//		list.PushBack("name" + strconv.Itoa(i))
//	}
//
//	//遍历链表
//	for e := list.Front(); e != nil; e = e.Next() {
//		fmt.Println(e.Value)
//	}
//
//	//链表长度
//	fmt.Println("len:", list.Len());
//	//链表清空
//	list.Init()
//	fmt.Println("len:", list.Len())
//}
