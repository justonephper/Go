1.ubuntu环境下安装go语言
```
  下载最新的二进制包，wget https://dl.google.com/go/go1.12.5.linux-amd64.tar.gz
```
  
  2.解压到/usr/local/go
  ```
  tar -C /usr/local/src -xzf go1.10.3.linux-amd64.tar.gz
  ```
  
  3.创建工作区
  ```
  mkdir /usr/local/go
  cd /usr/local/go
    bin : 生成的可执行文件的目录
    pkg: 编译生成的包的目标文件目录
    src : src下面的每个目录，就是一个包， 包内就是golang的源码文件
    
  ```
  
  4.添加命令行
  ```
  vim /etc/profile
    添加代码：
      export GOROOT=/usr/local/src/go
      export PATH=$PATH:$GOROOT/bin
      export GOPATH=/usr/local/go
      
    :wq! 退出
    source /etc/profile
    
  ```
    
  5. 测试：
  ```
    go version
    go env
  ```
