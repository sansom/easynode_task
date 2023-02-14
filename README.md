#easynode_task

easynode_task是easynode系统的基础和核心服务，是其他服务的运行的必要条件。
该服务负责任务产生、任务分发、系统监控等功能。

## Prerequisites
- go version >=1.16
- easynode_collect 服务已完成部署

## Building the source

(以linux系统为例)
- mkdir easynode & cd easynode
- git clone https://github.com/uduncloud/easynode_task.git
- cd easynode_task
- CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o easynode_task app.go
  (mac下编译linux程序为例，其他交叉编译的命令请自行搜索)

- ./easynode_task -config ./config.json

## config.json 详解


## usages

启动服务后，等待分配任务并执行