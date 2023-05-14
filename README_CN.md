# 麻雀 / Sparrow

<p style="text-align: center;">
  <a href="README.md">ENGLISH</a> | <a href="README_CN.md"  target="_blank">中文文档</a>
</p>

[![CodeQL](https://github.com/soulteary/sparrow/actions/workflows/codeql.yml/badge.svg)](https://github.com/soulteary/sparrow/actions/workflows/codeql.yml) [![Security Scan](https://github.com/soulteary/sparrow/actions/workflows/scan.yml/badge.svg)](https://github.com/soulteary/sparrow/actions/workflows/scan.yml) [![Release](https://github.com/soulteary/sparrow/actions/workflows/release.yaml/badge.svg)](https://github.com/soulteary/sparrow/actions/workflows/release.yaml) [![Go Report Card](https://goreportcard.com/badge/github.com/soulteary/sparrow)](https://goreportcard.com/report/github.com/soulteary/sparrow)

<img src=".github/logo.png" width="30%">

兼容 ChatGPT 接口风格的开源后端实现。

## 使用方法

服务可以独立运行，但是你如果想得到一个快速的效果，则需要使用 OpenAI 的客户端的请求方式，一个简单的方式，是使用 [soulteary/docker-chatgpt](https://github.com/soulteary/docker-chatgpt) 中的客户端。

你可以通过下面的命令，来直接启动一个容器版本的后端服务：

```bash
# download this project image from dockerhub
docker pull soulteary/sparrow
# refer to `docker-compose.yml` to add some environment variables you need
docker run -p 8091:8091 soulteary/sparrow
```

或者，下载项目中提供的 `docker-compose.yml` 示例配置文件，然后在配置文件所在的目录中执行下面的命令：

```bash
docker compose up -d
```

## 关于项目

- 如果 ChatGPT 生态中的内容，可以被连接或者被**任何其他服务**调用，会不会很有趣？

- 如果**开源生态中的任意产品和服务**功能，能够自由等接入 ChatGPT 或者类似交互模式的服务，会不会是一件有趣的事情？

- 如果我们能够将世界上最先进的各种服务，和充满自由的开源软件或者服务结合，尤其是那些没有提供 Chat OPS 交互能力的软件，会不会能够让很多事情变得轻松呢？

## 快速概览

**因为 “Talk is cheap”，所以，我会给你视频，当然，还有代码。**

- 2023.05.05 [在 ChatGPT 中玩 Midjourney](https://www.zhihu.com/pin/1637642465724325890)

无论用户是否使用插件，都应该可以在 ChatGPT 或类似的开源软件中自由使用其他在线服务。无论是商业服务还是免费的开源软件。

- 2023.05.01 [使用 ChatGPT 渲染长文内容](https://www.zhihu.com/pin/1636158221214887936)

如果我们要查找的内容，本身就是一个很长等内容，那么我们等客户端应该能够支持这么长内容的输出和展示，而不是硬生生的阶段。当然，这个是可选的，简洁有时候也很棒。

- 2023.03.05 [使用 ChatGPT 浏览网站并搜索游戏商品](https://www.zhihu.com/zvideo/1615679760738250752)

无论用户是否被 Waitlist 选中，通过 Chat Ops 的交互方式，访问世界并获取信息应该是最基本的权利。
