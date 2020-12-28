1.代码热更新功能实现，工具有很多（简述5中工具）,记录日期（2020-09-21）
```
  1.gin
    github地址:https://github.com/codegangsta/gin Star:3.5K
  
  2.realize
    github地址：https://github.com/oxequa/realize Star:3.8K
  
  3.Fresh
    github地址：https://github.com/gravityblast/fresh Star:2.8K
  
  4.bee
    github地址：https://github.com/beego/bee Satr:1.1K
  
  5.Air
    github地址：https://github.com/cosmtrek/air Star:2.0k
    
 ```
    
  2.由于最近正自使用gin-web框架编写项目，所以就是用了gin同名的热更新工具
  ```
  看下gin 的官方解释

  gin是一个简单的命令行实用程序，用于实时重新加载Go Web应用程序。 只需在您的应用程序目录中运行gin ，您的网络应用程序将以 gin 作为代理服务。
  当gin检测到有代码更改时，它会自动重新编译代码。 您的应用将在下次收到HTTP请求时重新启动。
  
  gin run main.go 
  监听已经启动，gin监听工具采用代理的工作模式，默认访问端口是3000，代理端口是3001，你需要设置一下代理端口即可
  
  方法1：直接修改gin run main.go中的源码，将代理端口改成8080（默认3001），采用硬编码的形式解决，显然这不是一个很完美的办法
  
  方法2：借助系统常量，BIN_APP_PORT,gin监听工具启动后，直接访问默认端口（3000）即可
      在 /etc/profile文件中加入 export BIN_APP_PORT=8080;source /etc/profile
      发现这个方法根本行不通，echo $BIN_APP_PORT，始终不输出值
      
  方法3：gin -h 使用帮助指令，查看具体的使用办法
  gin -p 3000 -a 8080 --all run
  
  options:
    -p 访问端口
    -a 代理端口
  
  访问 127.0.0.1:3000,随意修改代码，发现此时代码会自动重新部署
  ```
  
  completed!
      
  
  
  
