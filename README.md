# GoPkg

GoPkg 是一个为 Go 应用程序提供的实用工具集合，包含多种常用功能模块。

## 📖 简介

GoPkg 旨在提供一套完整、易用的 Go 语言工具包，帮助开发者快速构建高质量的应用程序。本项目包含多个独立的模块，涵盖了日常开发中常见的功能需求。

## ✨ 特性

- 🔒 **JWT** - JSON Web Token 实现
- 🔑 **签名** - 多种签名算法与验证
- 🚦 **限流** - 实用的速率限制功能
- 📱 **第三方 API** - 常用第三方 API 集成
- 📊 **数据结构** - 常用数据结构实现
- 📧 **邮件** - 邮件发送功能
- 🌐 **IP 定位** - IP 地址位置查询
- 🔍 **验证器** - 数据验证工具
- 🔄 **同步** - 并发控制工具
- 📂 **OSS** - 对象存储服务集成
- 🌈 **颜色** - 颜色处理工具
- 🌲 **树结构** - 树形数据结构操作
- 📄 **分页** - 分页功能实现
- 📊 **监控** - 系统监控工具
- 🔒 **权限控制** - 基于 Casbin 的权限控制

## 📦 安装

```bash
go get github.com/fzf-labs/gopkg
```

## 🚀 快速开始

以下是简单的使用示例：

```go
package main

import (
    "fmt"
    "github.com/fzf-labs/gopkg/jwt"
)

func main() {
    // 创建一个新的 JWT 令牌
    token, err := jwt.NewToken("your-secret-key", map[string]interface{}{
        "user_id": 123,
    }, 3600)
    if err != nil {
        panic(err)
    }
    
    fmt.Println("生成的令牌:", token)
    
    // 验证令牌
    claims, err := jwt.ParseToken(token, "your-secret-key")
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("解析的声明: %+v\n", claims)
}
```

## 📚 模块列表

- **algorithm**: 常用算法实现
- **binary**: 二进制数据处理
- **browser**: 浏览器检测工具
- **casbin**: 权限控制集成
- **color**: 颜色处理工具
- **datastructure**: 数据结构实现
- **ddm**: 分布式数据管理
- **email**: 电子邮件功能
- **encoding**: 编码/解码工具
- **goscraper**: 网页爬虫工具
- **iplocation**: IP 位置查询
- **jwt**: JWT 认证实现
- **monitor**: 系统监控工具
- **notes**: 注释与文档相关
- **oss**: 对象存储服务
- **page**: 分页功能
- **ratelimit**: 速率限制
- **reflection**: 反射相关工具
- **sign**: 签名算法实现
- **signature**: 签名验证
- **sync**: 同步与并发控制
- **third_api**: 第三方 API 集成
- **transport**: 传输相关工具
- **tree**: 树形结构操作
- **validatorx**: 验证工具扩展
- **version**: 版本管理

## 📄 许可证

该项目采用 [MIT 许可证](LICENSE)。

## 🤝 贡献

欢迎提交问题和拉取请求！在开始之前，请确保阅读 [贡献指南](CONTRIBUTING.md)。

## 🙏 致谢

感谢所有为本项目做出贡献的开发者！
