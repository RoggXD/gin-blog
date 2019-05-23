# Gin blog API

[![996.icu](https://img.shields.io/badge/link-996.icu-red.svg)](https://996.icu) [![LICENSE](https://img.shields.io/badge/license-Anti%20996-blue.svg)](https://github.com/996icu/996.ICU/blob/master/LICENSE)

Gin blog API 项目是基于 Golang Gin 框架的一套功能完整性能优良博客 API 服务，集博客基础操作、服务热更新、API 权限管理、API 可视化、图片上传下载、定时任务、系统日志采集等功能为一体。

本项目使用 [Glide](https://github.com/bumptech/glide) 管理项目依赖，Glide 非常好用，Golang 包管理工具大同小异具体选择看个人喜好了。

## 安装

先决条件：`Golang >= 1.8` `mysql` `redis`

```bash
$ go get github.com/RoggXD/gin-blog
```

## 试飞

### 数据库准备

首先请执行 **[blog.sql](./blog.sql)** 创建项目数据库和表。

### 项目基础配置

打开 `conf/app.ini` 配置项目配置。

### 起飞前最后的检查！

在试飞之前请检查项目目录 `runtime` 下是否有 `logs` 和 `upload/images` 目录，**这点非常重要**。

### 起飞！

```bash
$ cd $GOPATH/src/gin-blog

$ go run main.go 
```

> 项目包含热更新功能，修改配置后重启项目不会干扰已有链接。

## 获取 API 权限

获取 Token

```bash
http://host:8000/auth?username=test&password=test123456
```

## API 可视化

### 安装 swag

```bash
$ go get -u github.com/swaggo/swag/cmd/swag
```

若无法科学上网，需使用 gopm

```bash
gopm get -g -v github.com/swaggo/swag/cmd/swag

cd $GOPATH/src/github.com/swaggo/swag/cmd/swag

go install
```

### 生成

在项目根目录下，

```bash
swag init
```

会自动生成 `docs/`

```
docs/
└── docs.go
└── swagger
    └── swagger.json
    └── swagger.yaml
```

访问 `http://host:8000/swagger/index.html` 可视化查看 API 详情。

## 项目一级目录说明

```
gin-blog/
├── conf
├── middleware
├── models
├── routers
└── runtime
└── vendor
└── pkg
```

- conf：用于存储配置文件
- middleware：应用中间件
- models：应用数据库模型
- routers：路由逻辑处理
- runtime：应用运行时数据
- vendor：项目依赖
- pkg：第三方包

