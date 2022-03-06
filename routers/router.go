package routers

import (
	"net/http"

	"github.com/BM-laoli/go-gin-example/docs"
	"github.com/BM-laoli/go-gin-example/middleware/jwt"
	"github.com/BM-laoli/go-gin-example/pkg/export"
	"github.com/BM-laoli/go-gin-example/pkg/qrcode"
	"github.com/BM-laoli/go-gin-example/pkg/setting"
	"github.com/BM-laoli/go-gin-example/pkg/upload"
	"github.com/BM-laoli/go-gin-example/routers/api"
	v1 "github.com/BM-laoli/go-gin-example/routers/api/v1"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath())) // 路径静态资源访问支持 ，每个静态资源访问都需要独立去控制
	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))
	r.POST("/upload", api.UploadImage) // 上传图片

	// 获取文件
	r.POST("/auth", api.GetAuth)
	r.POST("/authRegister", api.Register)
	// 路由模块和分组 以及jwt验证中间价
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		// 标签
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		// 文章
		apiv1.POST("/articles/poster/generate", v1.GenerateArticlePoster)
		apiv1.GET("/articles", v1.GetArticles)
		apiv1.GET("/articles/:id", v1.GetArticle)
		apiv1.POST("/articles", v1.AddArticle)
		apiv1.PUT("/articles/:id", v1.EditArticle)
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)

		// 导出文件、导入标签、获取QR
		r.POST("/tags/export", v1.ExportTag)
		r.POST("/tags/import", v1.ImportTag)
	}

	return r
}
