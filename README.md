# Gin blog API

[![996.icu](https://img.shields.io/badge/link-996.icu-red.svg)](https://996.icu) [![LICENSE](https://img.shields.io/badge/license-Anti%20996-blue.svg)](https://github.com/996icu/996.ICU/blob/master/LICENSE)

先决条件：`Golang >= 1.8` `mysql` `redis`

本项目使用 [Glide](https://github.com/bumptech/glide) 管理项目依赖，Glide 非常好用，Golang 包管理工具大同小异具体选择看个人喜好了。

首先请执行 **[blog.sql](./blog.sql)** 创建项目数据库和表。

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

## 获取 API 权限

获取 Token

```bash
http://host:8000/auth?username=test&password=test123456
```

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

特点

- 热更新(需 `Glang >= 1.8`)
- API 可视化

