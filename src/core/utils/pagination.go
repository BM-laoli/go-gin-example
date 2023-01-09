package utils

import (
	configuration "github.com/BM-laoli/go-gin-example/src/core/configuration"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// 看看一共有多少条
func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * configuration.AppSetting.PageSize
	}

	return result
}
