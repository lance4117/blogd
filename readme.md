# blogd — A Simple Blog Blockchain (Cosmos SDK)

> 一个基于 **Cosmos SDK** 与 **CometBFT/Tendermint** 构建的去中心化博客区块链，支持链上发布、存储与查询帖子数据。项目由 **Ignite CLI** 脚手架初始化，支持 gRPC 与 REST 接口，提供完整的 CLI 体验。

---

## ✨ 特性

- ✅ 链上创建/更新/删除帖子（Msg 交易）
- 🔎 分页查询帖子、查询单帖、查询模块参数（Query 服务）
- 🌐 gRPC 与 REST API（由 `proto` 自动生成）
- 🧰 一键开发体验：`ignite chain serve`
- 🧪 包含 `x/blog` 模块、`proto`、`api` 以及二次开发所需的工程脚手架

> 快速开始命令与配置来自 Ignite 官方脚手架约定。详情见项目初始化页提示。  
> 参考：在仓库首页已有的 “Get started / Configure / Release / Install” 指引（例如 `ignite chain serve`）。
>

---

## 🧱 目录结构

```bash
.
├── app/               # 应用装配（编码、模块注册等）
├── x/blog/            # blog 模块：types、keeper、module.go
│   ├── keeper/
│   ├── types/
│   └── module.go
├── proto/             # Protobuf（tx.proto / query.proto 等）
├── api/blog/blog/     # 由 proto 生成的 gRPC/REST 适配层
├── cmd/blogd/         # 可执行入口（blogd）
├── config.yml         # Ignite 开发配置
├── Makefile
├── go.mod / go.sum
└── README.md

