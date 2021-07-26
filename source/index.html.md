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
---

# 介绍

[https://github.com/chyroc/lark](https://github.com/chyroc/lark) 的文档站点。

本 SDK 依靠自动读取 open.feishu.cn 网站的开发文档，解析文档结构，自动生成 SDK 内容。

所有内容全部依靠自动生成，可以及时更新。

# 安装

本 SDK 现在支持两个语言的版本，其他语言的版本正在生成过程中。

```go
go get github.com/chyroc/lark
```

```python
pip install pylark
```

# 初始化

Lark 开放接口支持多个场景的使用，包括：

- 常规 App
- 服务台
- 自定义机器人
- ISV App

不同的使用场景有不同的初始化方式

## 常规初始化

> 常规初始化

```go
file-embed: ./libs/code/1-init-normal.go
```

```python
import kittn

api = kittn.authorize('meowmeowmeow')
```

<aside class="success">
success
</aside>

<aside class="notice">
notice
</aside>

<aside class="warning">
warning
</aside>

## ISV 应用初始化

## 服务台应用初始化

