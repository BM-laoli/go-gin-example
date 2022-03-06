这个文章会有你很多值得参考的东西 - gin中文文档
<https://www.topgoer.cn/docs/ginkuangjia/ginkuangjia-1c50hfaag99k2>

煎鱼大神的博客
<https://eddycjy.com/posts/go/gin/2018-02-15-log/>

go-gin-example
├── conf
│   └── app.ini
├── docs                                // swag文档
├── middleware
│   └── jwt
│         └── jwt.go                  // jwt验证其中间价
├── models
│   └── article.go                    // 每一个模型 可以理解为DTO
│   └── auth.go
│   └── models.go
│   └── models.go                   // 主要的入口文件
│   └── tag.go
├── pkg
│   └── app                            // 收口App全局方法
│         └── form.go
│         └── request.go
│         └── response.go
│   └── e                                // 错误提示和错误代 码定义
│         └── cache.go
│         └── code.go
│         └── msg.go
│   └── export                          // 导入功能配置项
│   └── file                              // 文件判断公共方法
│   └── gredis                          // redis连接 配置和操作配置
│   └── logging                        // 日志记录器
│         └── file.go
│         └── log.go
│   └── qrcode                        // 二维码功能配置
│   └── setting                        // setting映射项
├── upload                              // image图片上传的方法

├── router
│   └── router.go
│   └── api
│         └── v1                          // v1 下的tag 和 article 路由
│               └── tag.go
│               └── article.go  
│         └── auth.go                  // 获取token 和上传文件服务
│         └── upload.go
├── runtime                             // runtime时的文件存储位置
│   └── export
│   └── fonts
│   └── logs
│   └── qrcode
│   └── upload
├── service
│   └── article_service           // 文章和 posts 海报生成的service
│   └── cace_service  // redis缓存相关的服务
│   └── tag_service              // tag服务
├── utils
│   └── jwt.go
│   └── md5.go
│   └── pagination
├── cron.go.test //  这原来是一个定时任务的脚本，但是由于是main包的所以我们先把它下架更改了
├── Dockerfile
└── go.mod
└── go.sum
└── main.go
└── README.md
