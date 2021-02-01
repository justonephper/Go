# 使用插件日志插件（github.com/sirupsen/logrus）

```
    日志是程序中必不可少的一个环节，由于Go语言内置的日志库功能比较简洁，我们在实际开发中通常会选择使用第三方的日志库来进行开发。
本文介绍了logrus这个日志库的基本使用。
```

## logrus介绍
> Logrus是Go（golang）的结构化logger，与标准库logger API完全兼容。
```
  它有以下特点：
      + 完全兼容标准日志库，拥有七种日志级别：Trace, Debug, Info, Warning, Error, Fataland Panic。
      + 可扩展的Hook机制，允许使用者通过Hook的方式将日志分发到任意地方，如本地文件系统，logstash，elasticsearch或者mq等，或者通过Hook定义日志内容和格式等
      + 可选的日志输出格式，内置了两种日志格式JSONFormater和TextFormatter，还可以自定义日志格式
      + Field机制，通过Filed机制进行结构化的日志记录
      + 线程安全
```

## 安装
```
    go get github.com/sirupsen/logrus
```

## 基本示例
> 使用Logrus最简单的方法是简单的包级导出日志程序:
```
    package main

    import (
       log "github.com/sirupsen/logrus"
    )

    func main() {
         log.WithFields(log.Fields{
           "animal": "dog",
        }).Info("一条舔狗出现了。")
    }
```

## 进阶示例
> 对于更高级的用法，例如在同一应用程序记录到多个位置，你还可以创建logrus Logger的实例：
```
    package main

    import (
      "os"
      "github.com/sirupsen/logrus"
    )

    // 创建一个新的logger实例。可以创建任意多个。
    var log = logrus.New()

    func main() {
      // 设置日志输出为os.Stdout
      log.Out = os.Stdout

      // 可以设置像文件等任意`io.Writer`类型作为日志输出
      // file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
      // if err == nil {
      //  log.Out = file
      // } else {
      //  log.Info("Failed to log to file, using default stderr")
      // }

      log.WithFields(logrus.Fields{
        "animal": "dog",
        "size":   10,
      }).Info("一群舔狗出现了。")
```

## 日志级别
> Logrus有七个日志级别：Trace, Debug, Info, Warning, Error, Fataland Panic。
```
    log.Trace("Something very low level.")
    log.Debug("Useful debugging information.")
    log.Info("Something noteworthy happened!")
    log.Warn("You should probably take a look at this.")
    log.Error("Something failed but I'm not quitting.")
    // 记完日志后会调用os.Exit(1) 
    log.Fatal("Bye.")
    // 记完日志后会调用 panic() 
    log.Panic("I'm bailing.")
```
> 设置日志级别
你可以在Logger上设置日志记录级别，然后它只会记录具有该级别或以上级别任何内容的条目：
```
  // 会记录info及以上级别 (warn, error, fatal, panic)
  log.SetLevel(log.InfoLevel)
```
如果你的程序支持debug或环境变量模式，设置log.Level = logrus.DebugLevel会很有帮助。

## 字段
Logrus鼓励通过日志字段进行谨慎的结构化日志记录，而不是冗长的、不可解析的错误消息。

例如，区别于使用log.Fatalf("Failed to send event %s to topic %s with key %d")，你应该使用如下方式记录更容易发现的内容:
```
    log.WithFields(log.Fields{
      "event": event,
      "topic": topic,
      "key": key,
    }).Fatal("Failed to send event")
```
WithFields的调用是可选的
