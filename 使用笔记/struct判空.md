# 结构体判空

```

  package main

  import (
      "fmt"
      "reflect"
  )

  type A struct{
      name string
      age int
  }

  func (a A) IsEmpty() bool {
      return reflect.DeepEqual(a, A{})
  }

  func main() {

      var a A

      if a == (A{}) {  // 括号不能去
          fmt.Println("a == A{} empty")
      }

      if a.IsEmpty() {
          fmt.Println("reflect deep is empty")
      }
  }
  
  总结：
    方法1：使用接口包裹数据结构的初始化值，与判空值进行比较
    方法2：借助反射包中DeepEqual方法，进行比较
    
    
    
```
