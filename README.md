# 说明

这个文章会有你很多值得参考的东西 - gin中文文档
<https://www.topgoer.cn/docs/ginkuangjia/ginkuangjia-1c50hfaag99k2>

煎鱼大神的博客
<https://eddycjy.com/posts/go/gin/2018-02-15-log/>

## 目录结构

```
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
```

## 遇到的一些问题

### db 的表的结构

> 下面是 建表的sql

```sql
CREATE TABLE `blog_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '标签名称',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章标签管理';

CREATE TABLE `blog_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `tag_id` int(10) unsigned DEFAULT '0' COMMENT '标签ID',
  `title` varchar(100) DEFAULT '' COMMENT '文章标题',
  `desc` varchar(255) DEFAULT '' COMMENT '简述',
  `content` text,
  `created_on` int(11) DEFAULT NULL,
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(255) DEFAULT '' COMMENT '修改人',
  `cover_image_url` varchar(255) DEFAULT '' COMMENT 'cover图片',
  `deleted_on` int(10) unsigned DEFAULT '0',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章管理';

CREATE TABLE `blog_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `blog`.`blog_auth` (`id`, `username`, `password`) VALUES (null, 'test', 'test123456');

```

### 关于redis的问题

> 由于我们使用了redis来缓存我们的 articles 数据，有时候会出现异常，耍不到数据，这时候你可先清除掉redis的数据 再做其它的 getlist 看看，它就可以正常的刷数据到redis 里面去了

## 23 年 1月分 重构的目录结构

> 我们的目标是尽力 像 spring 看齐，如果你喜欢这种风格的项目结构，那真是太好了，如果你不喜欢 没关系，毕竟go 非常的灵活 （🐶 毕竟这个领域go 才刚刚开始）

由于我们将会对我们的项目动 "大手术", 所以我们需要对项目结构进行调整，所以原来的一些引用路径发生了变化, 我们需要下面的命令

### 首先我们要做的就是迁移文件夹结构📁

```go

// 有些东西能够被 定义定义成 dto，比如中国 service 它就不能定义成dto，
// 换另一种理解，这个srvice 实际上你上一个 对象，它上面有方法，我们调用的时候只是
// 传递了一个 初始值，而后它会依据这些初始值去 调用后续的方法
type Tag struct {
 ID         int
 Name       string
 State      int
 CreatedBy  string
 ModifiedBy string
 PageNum    int
 PageSize   int
}

func (t *Tag) ExistByName() (bool, error) {
 return models.ExistTagByName(t.Name)
}
```
