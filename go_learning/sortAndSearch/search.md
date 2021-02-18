# 折半查找

>折半查找的前提是：数据有序

```
  package search

  import (
    "fmt"
    "testing"
  )

  func binarySearch(checkSlice *[]int, low, high, findVal int) int {
    if low > high {
      fmt.Println("not found")
      return -1
    }

    //找到中间下标
    mid := (low + high) / 2
    if findVal > (*checkSlice)[mid] {
      return binarySearch(checkSlice, mid+1, high, findVal)
    } else if findVal < (*checkSlice)[mid] {
      return binarySearch(checkSlice, low, mid-1, findVal)
    } else {
      //fmt.Println("找到了")
      return mid
    }
  }

  func TestBinarySearch(t *testing.T) {

    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    pos := binarySearch(&arr, 0, len(arr), 9)
    t.Log(pos)
  }

```
