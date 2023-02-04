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

# 自定义验证
>
> 我们一步到位直接 用interceptor 封装一下

1. 需要提前了解的知识点

```go
type PerRPCCredentials interface {
    GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error)
    RequireTransportSecurity() bool
}

// 这个是 gRPC 默认提供用于自定义认证的接口，它的作用是将所需的安全认证信息添加到每个 RPC 方法的上下文中
// GetRequestMetadata：获取当前请求认证所需的元数据（metadata）
// RequireTransportSecurity：是否需要基于 TLS 认证进行安全传输
```

2. 回到具体的操作上 我们还是cv一下 上一节的代码，然后做如下的修改

```go
// 第一 先改造 server

++++
const PORT = "9001"

// 实现 PerRPCCredentials start
type Auth struct {
 AppKey    string
 AppSecret string
}

func (a *Auth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
 return map[string]string{"app_key": a.AppKey, "app_secret": a.AppSecret}, nil
}

func (a *Auth) RequireTransportSecurity() bool {
 return true
}

func (a *Auth) Check(ctx context.Context) error {
 md, ok := metadata.FromIncomingContext(ctx)
 if !ok {
  return status.Errorf(codes.Unauthenticated, "自定义认证 Token 失败")
 }

 var (
  appKey    string
  appSecret string
 )
 if value, ok := md["app_key"]; ok {
  appKey = value[0]
 }
 if value, ok := md["app_secret"]; ok {
  appSecret = value[0]
 }

 if appKey != a.GetAppKey() || appSecret != a.GetAppSecret() {
  return status.Errorf(codes.Unauthenticated, "自定义认证 Token 无效")
 }

 return nil
}

type SearchService struct {
 pb.SearchServiceServer
 auth Auth
}

func (a *Auth) GetAppKey() string {
 return a.AppKey
}

func (a *Auth) GetAppSecret() string {
 return a.AppSecret
}


func CustomerAuth(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

 auth := Auth{
  AppKey:    "bmlaoli",
  AppSecret: "861795660",
 }

 //  if 如果头上有数据就成功
 if err := auth.Check(ctx); err != nil {
  log.Println("失败了")
  return nil, err
 }

// 一定要 handler 哈 要不然就没法next 下去了
 resp, err2 := handler(ctx, req)
 if err2 != nil {
  return nil, err2
 }
 return resp, err
}


// 实现 PerRPCCredentials end
++++
 opts := []grpc.ServerOption{
  grpc.Creds(c),
  grpc_middleware.WithUnaryServerChain(
   RecoveryInterceptor,
   LoggingInterceptor,
   CustomerAuth, // 把这个自定义的 验证器 加上就好了
  ),
 }
++++

// 对于客户端也是一样的 
// 实现 PerRPCCredentials start
type Auth struct {
 AppKey    string
 AppSecret string
}

func (a *Auth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
 return map[string]string{"app_key": a.AppKey, "app_secret": a.AppSecret}, nil
}

func (a *Auth) RequireTransportSecurity() bool {
 return true
}

// 实现 PerRPCCredentials end

func main() {
  ++++
 auth := Auth{
  AppKey:    "bmlaoli",
  AppSecret: "8617956603",
 }
  ++++
//直接加一个 WithPerRPCCredentials
 conn, err := grpc.Dial(":"+PORT, grpc.WithTransportCredentials(c), grpc.WithPerRPCCredentials(&auth))
  ++++
}

```

# gRPC Deadlines
>
> 主要就是设置一个 最终时间，到了就不让请求或者程序进行下去了，默认的grpc中 这个时间上有点长的（DEADLINE_EXCEEDED）。(煎鱼🐟的文章有点旧，但是并没过时 我们还是用他的文章来做)，我们直接看最新的 [官方文档](https://grpc.io/blog/deadlines/),  一个不错的博客的解析[link]( https://www.cnblogs.com/jiujuan/p/13499915.html )。

```go
// 设置一个deadline 主要是client 
++++

 ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Duration(5*time.Second)))
 defer cancel()

 //直接加一个 WithPerRPCCredentials
 conn, err := grpc.Dial(":"+PORT, grpc.WithTransportCredentials(c), grpc.WithPerRPCCredentials(&auth))

 if err != nil {
  log.Fatalf("grpc.Dial err: %v", err)
 }

 defer conn.Close()

 client := pb.NewSearchServiceClient(conn)
 resp, err := client.Search(ctx, &pb.SearchRequest{
  Request: "gRPC http",
 })

++++

// 检查 deadline （client 和service 都要检查 ）
++++Client
 if err != nil {
  statusErr, ok := status.FromError(err)
  if ok {
   if statusErr.Code() == codes.Canceled { // 检查 Canceled 字段
    log.Fatalln("client.Search err: deadline")
   }
  }

  log.Fatalf("client.Search err: %v", err)
  return
 }

 log.Printf("resp: %s", resp.GetResponse())
++++ service 
   if ctx.Err() == context.Canceled {
   return nil, status.Errorf(codes.Canceled, "SearchService.Search canceled")
  }

// 判断 deadline 之后的后续动作 （client 和 service ）这里不写了

// 补充一段 让service sleep 的代码，以达到“time out " 目的 
+++++
func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {

 // 休眠 service 也要做time out 检查 这个检查就 看 Canceled 字段
 for i := 0; i < 5; i++ {
  if ctx.Err() == context.Canceled {
   return nil, status.Errorf(codes.Canceled, "SearchService.Search canceled")
  }

  time.Sleep(1 * time.Second)
 }

 return &pb.SearchResponse{Response: r.GetRequest() + " Server!!!"}, nil
}

```

# 分布式链路追踪 gRPC + Opentracing + Zipkin

> 煎鱼的文档有点点的旧 版本多多少少也有点 问题，安装它的文档是运行不起来的； 我们直接去 看下面这些内容
<https://github.com/openzipkin/zipkin>，
<https://pkg.go.dev/github.com/openzipkin/zipkin-go#section-readme>
<https://pkg.go.dev/github.com/openzipkin-contrib/zipkin-go-opentracing>

知乎文章，掘金文章
<https://zhuanlan.zhihu.com/p/453190067>,
<https://juejin.cn/post/7051874138877198344>

1. docker 把 zipkin 平台运行起来 (过于简单这里不说了)

2. 安装必要的依赖 到我们的程序中, 与 zipkin 平台进行集成

> 重要的事情！！！ 一定一定要注意版本 v2！ 不是 v1 v1 你就400 了

依赖是

```go
 import (
    "github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
    "github.com/opentracing/opentracing-go"
    zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
    "github.com/openzipkin/zipkin-go"
    "github.com/openzipkin/zipkin-go/reporter"
    httpreport "github.com/openzipkin/zipkin-go/reporter/http"
    ++++
  )
```

代码是

```go
// ---------------------------------server

const (
 serviceName        = "simple_zipkin_server"
 zipkinAddr         = "http://127.0.0.1:9411/api/v2/spans"
 zipkinRecorderAddr = "127.0.0.1:9001"
)

// 创建一个zipkin追踪器
// url:http://localhost:9411/api/v2/spans
// serviceName:服务名，Endpoint标记
// hostPort：ip:port，Endpoint标记
func NewZipkinTracer(url, serviceName, hostPort string) (*zipkin.Tracer, reporter.Reporter, error) {

 // 初始化zipkin reporter
 // reporter可以有很多种，如：logReporter、httpReporter，这里我们只使用httpReporter将span报告给http服务，也就是zipkin的http后台
 r := httpreport.NewReporter(url)

 //创建一个endpoint，用来标识当前服务，服务名：服务地址和端口
 endpoint, err := zipkin.NewEndpoint(serviceName, hostPort)
 if err != nil {
  return nil, r, err
 }

 // 初始化追踪器 主要作用有解析span，解析上下文等
 tracer, err := zipkin.NewTracer(r, zipkin.WithLocalEndpoint(endpoint))
 if err != nil {
  return nil, r, err
 }

 return tracer, r, nil
}

// zipkin start
 tracer, r, err := NewZipkinTracer(zipkinAddr, serviceName, zipkinRecorderAddr)
 defer r.Close()
 if err != nil {
  log.Println("tracking-<", err)
  return
 }
 t := zipkinot.Wrap(tracer)
 opentracing.SetGlobalTracer(t)
 // zipkin end

 opts := []grpc.ServerOption{
  grpc.Creds(c),
  // grpc.StatsHandler(zipkingrpc.NewServerHandler(tracer)),
  grpc_middleware.WithUnaryServerChain(
   RecoveryInterceptor,
   LoggingInterceptor,
   otgrpc.OpenTracingServerInterceptor(t, otgrpc.LogPayloads()),
  ),
 }

// ---------------------------------client
 // zipkin start
 tracer, r, err := NewZipkinTracer(zipkinAddr, serviceName, zipkinRecorderAddr)
 defer r.Close()
 if err != nil {
  log.Println("tracer %d", err)
  return
 }
 t := zipkinot.Wrap(tracer)
 opentracing.SetGlobalTracer(t)
 // zipkin end

   conn, err := grpc.Dial(":"+PORT, grpc.WithTransportCredentials(c),
  grpc.WithUnaryInterceptor(
   otgrpc.OpenTracingClientInterceptor(t, otgrpc.LogPayloads()),
  ),
 )

// ---------------------------------其它说明
// 特殊说明 实际上我们也 可以用官方 zipkin-go 文档中提到的 
 opts := []grpc.ServerOption{
  grpc.Creds(c),
  grpc.StatsHandler(zipkingrpc.NewServerHandler(tracer)),
  grpc_middleware.WithUnaryServerChain(
   RecoveryInterceptor,
   LoggingInterceptor,
  ),
 }

   conn, err := grpc.Dial(":"+PORT,
     grpc.WithTransportCredentials(c),
     grpc.WithStatsHandler(zipkingrpc.NewClientHandler(tracer)) )
```
