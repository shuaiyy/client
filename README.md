# 项目组织

```
easyai-core 同步自后端项目的接口sdk，请勿直接修改
seelie
├── README.md
├── client： 向服务端发起请求的client
├── command： cli的命令/子命令解析; 应包含demo case、help docs、url link等，提高工具的易用性
├── root.go： cli执行入口，`rootCommand.Execute()`
├── run：调用client，执行具体的service逻辑
└── utils：辅助工具库
```

调用链路

cmd/main.go 
  internal/cli root.go --> command --> run --> SeelieClient -http-> Seelie API Server

## 构建

```
cd seelie
make cli
```


# 长期

命令帮助文档的文案优化

# TASKS

+ 计划要做的功能，请阅读[tasks.md](./TASKS.md)