# 单元测试

## 格式要求
```
    1. 文件名格式：*_test.go
       例如： simple_test.go
       
    2. 方法名称：Test*，例如：
      func TestAdd(t *testing.T){
        num := 100
        //普通打印
        t.log(num)
        
        //格式化输出
        t.logf("%#v\n",num)
      }
```

## 执行指令
```
    go test -v **_test.go
    
    -v 展示详情的每个测试方法的执行结果，不加，只显示总的测试结果
```

## 测试用例与测试源码同文件，导致方法 undefined
```
    go test -v *_test.go *.go
    
    例如：
      go test -v simple_test.go simple.go
      后面跟上依赖的源码文件，就不会出现 undefined问题
```
