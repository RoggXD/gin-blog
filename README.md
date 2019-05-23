# Gin blog API

[![996.icu](https://img.shields.io/badge/link-996.icu-red.svg)](https://996.icu) [![LICENSE](https://img.shields.io/badge/license-Anti%20996-blue.svg)](https://github.com/996icu/996.ICU/blob/master/LICENSE)

本项目使用 [Glide](https://github.com/bumptech/glide) 管理项目依赖，Glide 非常好用，Golang 包管理工具大同小异具体选择看个人喜好了。

首先请执行 **[blog.sql](./blog.sql)** 创建项目数据库和表。

## 项目一级目录说明

``` sh
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