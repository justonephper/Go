# validate常用语法


## 校验请求参数

>gin框架使用[github.com/go-playground/validator](github.com/go-playground/validator)进行参数校验，目前已经支持[github.com/go-playground/validator/v10](github.com/go-playground/validator/v10)了，我们需要在定义结构体时使用 binding tag标识相关校验规则，
可以查看[validator文档](https://pkg.go.dev/github.com/go-playground/validator#hdr-Baked_In_Validators_and_Tags)查看支持的所有 tag。

```
    
    举例1：
    
    添加博客参数
    type BlogAddParams struct {
      Name    string `json:"name" binding:"required,Maximum=64"`
      Title   string `json:"title" binding:"required,Maximum=128"`
      Content string `json:"content" binding:"required"`
    }
    
    知识点：
      1. json可以给参数取别名
      2. binding语法，支持参数校验，binding后面多个值之间使用逗号 ","
      
      
   举例2：
   
   列表接口参数
   // Paging common input parameter structure
    type PageInfo struct {
        Page     int    `json:"page" form:"page" binding:"required"`
        PageSize int    `json:"page_size" form:"page_size" binding:"required"`
        Q        string `json:"q" form:"q"`
        OrderKey string `json:"order_key" form:"order_key"`
        Desc     bool   `json:"desc" form:"desc"`
    }
    
    知识点：
        1. page,page_size,q,order,desc等字段是通过url形式传输（xxx.com?page=1&page_size=6&q=健康&order_key=created_at&desc=true）
        2. 结构体中，需要声明form的tag才能解析上述参数
  
```

## 模型数据结构

> [gorm文档地址](https://gorm.io/zh_CN/docs/)
```
    举例：
    
    type Blog struct {
      ModelId
      Name    string `json:"name" gorm:"default:'';size:64;comment:'博客名称'"`
      Title   string `json:"title" gorm:"default:'';size:128;comment:'标题'"`
      Content string `json:"content" gorm:"type:text;not null;comment:'内容'"`
      ModelTime
    }
    
    知识点：
      1. json为字段起别名，json输出时自动转化
      2. gorm为数据迁移时，字段的要求,多个规则之间使用分号，";"
        常用：
          2.1 default:''          设置字段的默认值
          2.2 not null            字段不能为null
          2.3 size 128            设置字段的长度
          2.4 comment:"字段备注"   字段注释
          
          或者使用：
          type:varchar(64);
          
```
