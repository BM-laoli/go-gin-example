# 环境说明

> 首先是安装，protobuf 的过程和配置 本地开发环境变量

1. 安装

我是mac电脑，直接安装就好啦

```shell
$ brew install protobuf 
# 如果您遇见了一些 安装的报错，比如安了一般，还剩一半，这个时候，请有可能找不到对应的tackid的文件，你需要清理一下
# brew 的缓冲 然后再重新 brew install 

Warning: Bottle missing, falling back to the default domain...
==> Downloading https://ghcr.io/v2/homebrew/core/protobuf/manifests/3.17.3
Already downloaded: /Users/administrator/Library/Caches/Homebrew/downloads/6d20d704ec9f8c1fe7646d23a2e0e96e28c081140c98ef75911b36eae4e91f6f--protobuf-3.17.3.bottle_manifest.json
==> Downloading https://ghcr.io/v2/homebrew/core/protobuf/blobs/sha256:ef7a56961
==> Downloading from https://pkg-containers.githubusercontent.com/ghcr1/blobs/sh
######################################################################## 100.0%
==> Installing dependencies for protobuf: six
==> Installing protobuf dependency: six
==> Pouring six-1.16.0_1.arm64_big_sur.bottle.tar.gz
Error: No such file or directory @ rb_sysopen - /Users/administrator/Library/Caches/Homebrew/downloads/b2d32282c41ba4907220501baebea0c80fcc339a9b91684245d54261c17360b0--six-1.16.0_1.arm64_big_sur.bottle.tar.gz

# 比如现在就是 找不到 /Users/administrator/Library/Caches下的这个id号的问题，你可以先把 Caches 下都清理一遍，在重新安装
```

2. 关于go升级

```shell
brew upgrade go
```

3. 关于 protoc xxx指令错误

```shell
$ protoc --go-grpc_out= ./search.proto 
protoc-gen-go-grpc: program not found or is not executable
Please specify a program using absolute path or make sure the program is available in your PATH system variable
--go-grpc_out: protoc-gen-go-grpc: Plugin failed with status code 1.

```

首先先安装确保 一些正确的文件 （在你的项目root下）

```shell
$ go get -u google.golang.org/grpc

# 任选其一如果有问题
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
$ go get -u github.com/golang/protobuf/protoc-gen-go

```

然后 确保你的go env 的bin 执行是有关联的

```shell
$ go env 
# 先看看 go path 和 go bin 是否正确，如果不，请去修改,我这里是 zshrc ， 加入下面的内容，然后更新就好啦 source ~/.zshrc      ,最后去 go env 检查一下看看是否已经配置完，如果是那么久完成啦
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOPATH/bin

```

4. 依据协议 来转译 的时候，失败的问题

你需要用自己的包，比如 我自己的 github.com/BM-laoli/go-gin-example/proto

```go
syntax = "proto3";

package proto;

option go_package ="github.com/BM-laoli/go-gin-example/proto";

service SearchService {
    rpc Search(SearchRequest) returns (SearchResponse) {}
}

message SearchRequest {
    string request = 1;
}

message SearchResponse {
    string response = 1;
}

```

第二，你在 protoc 🔧, (执行的文件夹在项目的根目录go-gin-exampale)

```shell
$ protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/search.proto
```

这个命令完之后回生产两个 文件，具体请看官方文档 [go-rpc官方文档](https://grpc.io/docs/languages/go/quickstart/)

第三，你需要 实现它

```go
type SearchService struct {
 pb.SearchServiceServer
}
```
