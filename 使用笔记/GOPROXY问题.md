  设置代理常量：
  ```
    go env -w GOPROXY=https://goproxy.cn,direct // 使用七牛云的
  ```
  
  异常：warning: go env -w GOPROXY=... does not override conflicting OS environment variable
  
  删除常量
  ```
  unset GOPROXY
  ```

  查看常量：
  ```
  go env GOPROXY
  ```
  结果：常量生效
