# gin-blog

本项目使用 [Glide](https://github.com/bumptech/glide) 管理项目依赖，Glide 非常好用，Golang 包管理工具大同小异具体选择看个人喜好了。

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