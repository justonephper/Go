# gin框架支持自动迁移，我们只要编写初始化表的文件就可以

```
  例如：
  
  type Student struct {
    name      string
    password  string
  }
  
  //自动迁移
  gorm.DB.AutoMigrate(Student{})
  ...
  
```
