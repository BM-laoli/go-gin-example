package jwt

import (
	"net/http"
	"strings"
	"time"

	"github.com/BM-laoli/go-gin-example/pkg/e"
	"github.com/BM-laoli/go-gin-example/pkg/logging"
	util "github.com/BM-laoli/go-gin-example/utils"
	"github.com/gin-gonic/gin"
)

// 中间件 用于处理和检测token 是否正确
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		var token string

		code = e.SUCCESS

		var orgToken = c.Request.Header.Get("Authorization")
		if orgToken == "" {
			code = e.INVALID_PARAMS
		} else {
			token = strings.Fields(orgToken)[1]
		}

		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			logging.Error(err)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		// 成功之后 调用next 说明通过了这个中间件 和node 比较类似
		c.Next()
	}
}
