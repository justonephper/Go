# reflect
        反射是用程序检查其所拥有的结构，尤其是类型的一种能力；这是元编程的一种形式。反射可以在运行时检查类型和变量，例如它的大小、方法和动态
    的调用这些方法。这对于没有源代码的包尤其有用。这是一个强大的工具，除非真得有必要，否则应当避免使用或小心使用
    
## 变量的最基本信息就是类型和值：
        反射包的Type用来表示一个Go类型，反射包的Value为Go值提供了反射接口。
        
        两个简单的函数,reflect.TypeOf和reflect.ValueOf，返回被检查对象的类型和值。例如，x被定义为：var x float64 = 3.4,那么
    reflect.TypeOf(x)返回float64，reflect.ValueOf(x)返回3.4
    
### 获取类型信息和值信息
  
  ```
      type := reflect.TypeOf(num)
      value := reflect.ValueOf(num)
  ```
  
### 类型信息操作
  
  ```
      typeOf := relfect.TypeOf(num)
      
  ```
