## socket定义

  ```
  socket是BSD UNIX的进程通信机制，通常称为“套接字”，用于描述IP和端口，是一个通信链的句柄。socket可以理解为TCP/IP的API,它定义了一组函数，我们可以用它开发TCP/IP上的网络程序。
  电脑上通常通过socket向网络发起请求和应答网络上的请求。
  ```
  
## socket在网络中的位置
  ```
  Socket是应用层与TCP/IP协议族通信的中间软件抽象层。在设计模式中，Socket其实就是一个门面模式，它把复杂的TCP/IP协议族隐藏在Socket后面，对用户来说只需要调用Socket规定的相关函数，
  让Socket去组织符合指定的协议数据然后进行通信
  ```
