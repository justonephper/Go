1.使用go的模块工具
  ```
  模块代理网址：https://goproxy.cn/
  
  前置条件：1.GO的语言版本大于1.13，肯定好用
  
  设置环境变量（windows）： 
             go env -w GO111MODULE=on
             go env -w GOPROXY=https://goproxy.cn,direct
```
             
2.使用mod模块：
```
  go mod init api.fitness.com   
  生成go.mod文件
```

3.测试一下：
 ```
  go get -u github.com/gin-gonic/gin
  速度特别快，点赞
 ```

4.执行一下dity
```
  检查一下依赖，把没用的代码包删除，下载需要的代码包
```

Q1:下载工具包，使用指令 go get -u github.com/Unknow/com

结果提示：
go get: github.com/Unknwon/com@v1.0.1: parsing go.mod:module declares its path as: github.com/unknwon/com but was required as: github.com/Unknwon/com
  ```
  解决方案：
  1.1 查资料得知，使用go-mod 引入某个包出错例,需要在go.mod文件做处理
  
  1.2 格式为:
    replace module declares its path as:(后边那部分) => but was required as:(后边那部分)
  此处指令：
  replace github.com/unknwon/com github.com/unknwon/com v1.0.1
  
  1.3 发现，go.mod文件还是报红(编辑器不能正确索引包文件)，在命令行执行
  go mod dity
  
  执行后，go.mod文件能正确识别包文件，问题解决
  ```
  
  
