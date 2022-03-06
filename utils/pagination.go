package util

import (
	"github.com/BM-laoli/go-gin-example/pkg/setting"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// 看看一共有多少条
func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.AppSetting.PageSize
	}

	return result
}
