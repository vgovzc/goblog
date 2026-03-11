# goblog

a blog backend by golang

1.创建目录结构，并初始化

```
goblog
  |-config
  |-controller
  |-middleware
  |-model
  |-service
  |-util

go mod init goblog

```

2.添加引用

```
# 1.9.1 是稳定版，兼容 1.24
go get -u github.com/gin-gonic/gin@v1.9.1 
# 指定 GORM 版本（v1.25.4 是兼容 1.24 的稳定版）
go get -u gorm.io/gorm@v1.25.4
go get -u gorm.io/driver/mysql@v1.5.2
# 指定 Viper版本
go get -u github.com/spf13/viper@v1.18.2
```

3.
