## miniweb 项目

miniweb（微型博客） 是一个 Go 语言入门项目，用来完成用户注册、博客创建等业务。miniweb 入门但不简单：

- **入门：** 适用于刚学习完 Go 基础语法，零项目开发经验的 Go 开发者；
- **不简单：** 本项目来自于一线企业的大型线上项目，项目本身是一个企业级的项目，学习完之后，完全可以用来作为企业级项目的开发脚手架。

miniweb 实现了以下 2 类功能：
- **用户管理：** 支持 用户注册、用户登录、获取用户列表、获取用户详情、更新用户信息、修改用户密码、注销用户 7 种用户操作；
- **博客管理：** 支持 创建博客、获取博客列表、获取博客详情、更新博客内容、删除博客、批量删除博客 6 种博客操作。

**本项目适合人群**

- 刚学习完 Go 基础语法，想快速学习，以参与公司 Go 语言开发工作的开发者；
- 掌握 Go 基础语法，零 Go 应用开发经验，想通过完整的实战，快速、系统的学习 Go 开发的开发者；
- 有意从事 Go 应用开发，但尚未入门或入门尚浅的开发者；
- 有过 Go 应用开发经验，但想了解某一部分开发方法的开发者。


## Features

- 使用了简洁架构；
- 使用众多常用的 Go 包：gorm, casbin, govalidator, jwt-go, gin, cobra, viper, pflag, zap, pprof, grpc, protobuf 等；
- 规范的目录结构，使用 [project-layout](https://github.com/golang-standards/project-layout) 目录规范；
- 具备认证(JWT)和授权功能(casbin)；
- 独立设计的 log 包、error 包；
- 使用高质量的 Makefile 管理项目；
- 静态代码检查；
- 带有单元测试、性能测试、模糊测试、Mock 测试测试案例；
- 丰富的 Web 功能（调用链、优雅关停、中间件、跨域、异常恢复等）；
  - HTTP、HTTPS、gRPC 服务器实现；
  - JSON、Protobuf 数据交换格式实现；
- 项目遵循众多开发规范：代码规范、版本规范、接口规范、日志规范、错误规范、提交规范等；
- 访问 MySQL 编程实现；
- 实现的业务功能：用户管理、博客管理；
- RESTful API 设计规范；
- OpenAPI 3.0/Swagger 2.0 API 文档；
- 配套有高质量的掘金课程；


## Installation

```bash
$ git clone https://github.com/marmotedu/miniweb.git
$ go work use miniweb # 如果 Go 版本 > 1.18
$ cd miniweb
$ make # 编译源码
```

## Documentation

- [用户手册](./docs/guide/zh-CN/README.md)
- [开发手册](./docs/devel/zh-CN/README.md)

## Feedback

如果您有任何反馈，请通过 `18550039021@163.com` 与我联系。

## Contributing

欢迎你贡献代码和 Star !

### 开发规范

本项目遵循以下开发规范：[miniweb 项目开发规范](./docs/devel/zh-CN/conversions/README.md)。

## Authors

- [@莫维龙](https://www.github.com/bugsmo)


## License

[MIT](https://choosealicense.com/licenses/mit/)

## Related
