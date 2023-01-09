package routers

import (
	"net/http"

	auth "github.com/BM-laoli/go-gin-example/src/core/auth"
	configuration "github.com/BM-laoli/go-gin-example/src/core/configuration"
	core_upload "github.com/BM-laoli/go-gin-example/src/core/upload"
	middleware_jwt "github.com/BM-laoli/go-gin-example/src/middleware/jwt"
	"github.com/BM-laoli/go-gin-example/src/modules/article"
	tag "github.com/BM-laoli/go-gin-example/src/modules/tag"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(configuration.ServerSetting.RunMode)

	r.POST("/auth", auth.GetAuth)
	r.POST("/authRegister", auth.Register)

	// 上传文件
	r.POST("/upload", core_upload.UploadImage)
	r.StaticFS("/upload/images", http.Dir(core_upload.GetImageFullPath())) // 路径静态资源访问支持 ，每个静态资源访问都需要独立去控制

	// 路由模块和分组 以及jwt验证中间价
	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware_jwt.JWT())
	{
		// 标签
		apiv1.GET("/tags", tag.GetTags)
		apiv1.POST("/tags", tag.AddTag)
		apiv1.PUT("/tags", tag.EditTag)
		apiv1.DELETE("/tags/:id", tag.DeleteTag)

		// 文章
		apiv1.GET("/articles", article.GetArticles)
		apiv1.GET("/articles/:id", article.GetArticle)
		apiv1.POST("/articles", article.AddArticle)
		apiv1.PUT("/articles", article.EditArticle)
		apiv1.DELETE("/articles/:id", article.DeleteArticle)
	}

	return r
}
