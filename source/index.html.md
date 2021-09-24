---
title: Lark SDK

language_tabs: # must be one of https://git.io/vQNgJ

- go
- python

toc_footers:

- <a href='https://github.com/chyroc/lark'>https://github.com/chyroc/lark</a>
- <a href='https://github.com/slatedocs/slate'>Documentation Powered by Slate</a>

search: true

code_clipboard: true

includes:
- docs.md
---

# 介绍

[https://github.com/chyroc/lark](https://github.com/chyroc/lark) 的文档站点。

本 SDK 依靠自动读取 open.feishu.cn 网站的开发文档，解析文档结构，自动生成 SDK 内容。

所有内容全部依靠自动生成，可以及时更新。

# 安装

```go
go get github.com/chyroc/lark
```

```python
pip install pylark
```

本 SDK 现在支持两个语言的版本，其他语言的版本正在生成过程中。

# 初始化

Lark 开放接口支持多个场景的使用，包括：

- 常规 App
- 服务台
- 自定义机器人
- ISV App

不同的使用场景有不同的初始化方式

## 常规应用的初始化

> 常规应用的初始化

```go
file-embed: ./libs/code/1-init-normal.go
```

```python
file-embed: ./libs/code/1-init-normal.py
```

常规初始化只需要两个参数：app-id 和 app-secret，可以处理大部分场景。

## 回调应用的初始化

> 回调应用的初始化

```go
file-embed: ./libs/code/2-init-encrypt.go
```

file-embed: ./libs/code/2-init-encrypt.py
todo
```

处理回调的初始化，除了需要两个参数：app-id 和 app-secret 外，还需要 encryptKey 和 verificationToken 。

verificationToken 参数用于校验回调请求是否是合法的。

encryptKey 参数是可选的， 当你的应用启用了回调加密的时候，请务必填写本参数。

## 自定义机器人的初始化

> 自定义机器人的初始化

```go
file-embed: ./libs/code/3-init-custom-bot.go
```

```python
file-embed: ./libs/code/3-init-custom-bot.py
```

自定义机器人不需要单独创建应用，也不需要申请权限，只需要在群聊中添加自定义机器人即可。

自定义机器人也有一些限制，即只能发消息，也没办法处理回调。

在创建自定义机器人的时候，会给你一个回调链接 customBotWebHookURL 和一个秘钥 customBotSecret，填入初始化自定义机器人的参数即可。

## 服务台应用的初始化

> 服务台应用的初始化

```go
file-embed: ./libs/code/4-init-helpdesk.go
```

```python
file-embed: ./libs/code/4-init-helpdesk.py
```

服务台应用的初始化，除了需要两个参数：app-id 和 app-secret 外，还需要 helpdeskID 和 helpdeskToken，这两个参数可以在服务台的后台查看。

