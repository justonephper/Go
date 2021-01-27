# 1. reflect定义
        反射是用程序检查其所拥有的结构，尤其是类型的一种能力；这是元编程的一种形式。反射可以在运行时检查类型和变量，例如它的大小、方法和动态
    的调用这些方法。这对于没有源代码的包尤其有用。这是一个强大的工具，除非真得有必要，否则应当避免使用或小心使用
    
## 2. 变量的最基本信息就是类型和值：
        反射包的Type用来表示一个Go类型，反射包的Value为Go值提供了反射接口。
        
        两个简单的函数,reflect.TypeOf和reflect.ValueOf，返回被检查对象的类型和值。例如，x被定义为：var x float64 = 3.4,那么
    reflect.TypeOf(x)返回float64，reflect.ValueOf(x)返回3.4
    
### 3.1 获取类型信息和值信息
       type := reflect.TypeOf(num)
       value := reflect.ValueOf(num)
  
### 3.2 类型信息操作

      typeOf := relfect.TypeOf(num)

      //kind信息
      typeOf.Kind()

      //变量名称
      name := typeOf.Name()

      //数据是结构体时，查询字段数量
      for i := 0; i < typeOf.NumField(); i++ {}
  
  ### 3.3 值信息操作
  
      valueOf := reflect.ValueOf(num)
                
      //查询kind信息
      valueOf.Type().Kind()           // valueOf.Type() == typeOf
      typeOf.Kind() == valueOf.Kind()
                
      //接口原值=》接口值
      valueOf.interface()   
        
### 3.4 反射获取结构的方法

	type Student struct {
	    Name string
	    Age  int
	}

	func (stu *Student)SetName(name string) bool {
	    stu.Name = name
	    return true
	}
                
	func TestReflect(t *testing.T){
	    var stu Student {
		"haoge",
		25,
	    }
	    valueOf := reflect.ValueOf(&stu)

	    //第一种，访问下标的方式
	    valueOf.Method(0).Call(nil)

	    //第二种，指定方法名的方式
	    valueOf.MethodByName("SetName").Call(nil)
	}
        
### 3.5 遍历结构体

	type Student struct {
	    Name        string
	    Age         int
	    Score       string
	    Address     string
	}

	func TestReflect(t *testing>T) {
	    stu := Student {
		"haoge",
		23,
		100,
		"beijingshi",
	    }

	    valueOf := reflect.ValueOf(stu)
	    typeOf := reflect.TypeOf(stu)

	    name := typeOf.Name()
	    name = strings.ToLower(name)

	    query := fmt.Speinrf("insert into %s (\"name\",\"age\",\"score\",\"address\") values(",name)

	    for i:=0;i<valueOf.NumField();i++ {
		val := valueOf.Field(i)
		if i == 0 {
		    switch v.Type().Kind() {
			case reflect.Int:
			    query = fmt.Sprintf("%s%v", query, val)
			case reflect.String:
			    query = fmt.Sprintf("%s\"%v\"", query, val)
		    }
		} else {
		    switch val.Type().Kind() {
			case reflect.Int:
			    query = fmt.Sprintf("%s,%v", query, val)
			case reflect.String:
			    query = fmt.Sprintf("%s,\"%v\"", query, val)
		    }
		}
	    }

	    query = fmt.Sprintf("%s);", query)
	    t.Log(query)
	}
