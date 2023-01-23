> 这个分支下，我们会在主线上和煎鱼大神的go-grpc保持一致，然后会新开一些 小店，参照最新的文档和库(2023.1.12)来做补充说明和记录📝，具体的内容，和细枝末节，这里不会过多展开，煎鱼的blog 说的也比较的完善了。

# 环境说明(校准文档1)

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

# gRPC Streaming, Client and Server

> 这里主要说吗，三种不同的数据传输方式比，一方用流 一方不用，或双方都用

1. 三种传输方式

```proto
service StreamService {
// List：服务器端流式 RPC
// Record：客户端流式 RPC
// Route：双向流式 RPC
    rpc List(StreamRequest) returns (stream StreamResponse) {};

    rpc Record(stream StreamRequest) returns (StreamResponse) {};

    rpc Route(stream StreamRequest) returns (stream StreamResponse) {};
}
```

本章节，没有说明其他要说明的，就是简单的说一些三种stream 传输的简单使用，以及它们的使用场景

# 基于CA的TLS证书

> 前面的set debug 怎么能上生产呢？这是不可以的，所以我们需要加CA, 由于历史问题 ，以及煎鱼大神的blog 上这篇文章发布在2018年，现在我写这篇文章，已经上2023年，我们需要对照 grpc 官方来看, grpc 官方 + 煎鱼

这里有篇 [CSDN](https://blog.csdn.net/weixin_41335923/article/details/124250753)的 解读，还算不错的

grpc 官方的一个 ca 脚本生成器，我直接拿过来了 （注意你需要cv它们的 .conf 文件，并且你要有一定的知识，对你TAL，可以参考下面的一片文章）

[OpenSSL openssl.cnf](https://blog.csdn.net/wzfgd/article/details/109805158)

[create_ca.sh](https://github.com/grpc/grpc-go/blob/master/examples/data/x509/create.sh)

[这篇文章有错误，在实践的时候代码对不上的，也不知道它怎么敢发，但是理论内容可以看看](https://blog.csdn.net/scarificed/article/details/126184422)

grpc 官方 例子，[grcp-官方TSL例子](https://github.com/grpc/grpc-go/tree/master/examples/features/encryption)。里面的mTSL 就是于我们对应的内容

# Unary and Stream interceptor 拦截器
>
> 这个东西相对的简单，没有什么坑，这里就不展开说明了，请直接去看 官方文档，和煎鱼的文章就好了
[jainyuBlog](https://eddycjy.com/posts/go/grpc/2018-10-10-interceptor/),  [go-grpc-middleware](https://github.com/grpc-ecosystem/go-grpc-middleware)

多说一句 interceptor 和 middleware 的关系，我们可以用 Nest的工程来举例子，这里有一个我认为还是不错的文章 供你参考 [Nest项目结构中的各个部分](https://blog.csdn.net/lxy869718069/article/details/103960790)，当然我认为这不是绝对的顺序，有时候它取决你的 设计的，假设你的设计中  你把 interceptor 和 middleware 换了个位置。那么这个时候应该与你自己的有关

# 同时提供http/https 服务

> 核心步骤：判断 -> 转发 -> 响应。(注意 这只是一个简单的例子，并不完全符合生产标准)

这了和煎鱼大神的博客有点点的区别，我们并没有提取一些公共的ca 工具方法，我们直接cv 了上一个commit 中的  server / client ，然后做了下面的修改

```go

//  这是一个 handle 请求处理函数 当有http请求来的时候会走到这里
func GetHTTPServeMux() *http.ServeMux {
 mux := http.NewServeMux()
 mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("bm-laoli: go-inxxxxx"))
 })
 return mux
}

// 在原来的东西里面，判断一下就好了

// 去掉 
 // lis, err := net.Listen("tcp", ":"+PORT)
 // if err != nil {
 //  log.Fatalf("net.Listen err: %v", err)
 // }
 // server.Serve(lis)
 
//  缓成下面的内容 主要是直接用http TLS 煎听 请求，然后分别转到grpc 和 http 请求上

 http.ListenAndServeTLS(":"+PORT,
  certFile,
  keyFile,
  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
   if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
    server.ServeHTTP(w, r)
   } else {
    mux.ServeHTTP(w, r)
   }

   return
  }),
 )
```
