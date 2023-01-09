package middleware_jwt

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	constants "github.com/BM-laoli/go-gin-example/src/core/constants"
	core_error "github.com/BM-laoli/go-gin-example/src/core/error"
	log "github.com/BM-laoli/go-gin-example/src/core/log"
	core_redis "github.com/BM-laoli/go-gin-example/src/core/redis"
	"github.com/gin-gonic/gin"

	util "github.com/BM-laoli/go-gin-example/src/core/auth"

	error3 "github.com/BM-laoli/go-gin-example/src/core/error"
)

// 中间件 用于处理和检测token 是否正确
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		var token string

		code = constants.SUCCESS

		var orgToken = c.Request.Header.Get("Authorization")
		if orgToken == "" {
			code = constants.INVALID_PARAMS
		} else {
			token = strings.Fields(orgToken)[1]
		}

		if token == "" {
			code = constants.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = constants.ERROR_AUTH_CHECK_TOKEN_FAIL
				c.JSON(http.StatusUnauthorized, gin.H{
					"code": code,
					"msg":  core_error.GetMsg(code),
					"data": data,
				})
				c.Abort()
				return
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = constants.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				c.JSON(http.StatusUnauthorized, gin.H{
					"code": code,
					"msg":  core_error.GetMsg(code),
					"data": data,
				})
				c.Abort()
				return
			}

			// 看看redis 若有，获取看看是否== 如果没有 这个key 请登陆
			// 如果有这个key 请对比
			// 如果对比不成功 说明异地登陆了

			if !core_redis.Exists(claims.Uid) {

				c.JSON(http.StatusUnauthorized, gin.H{
					"code": code,
					"msg":  "redis中无数据",
					"data": data,
				})
				c.Abort()
				return
			}

			var cacheToken = ""
			data, err2 := core_redis.Get(claims.Uid)
			if err2 != nil {
				log.Error(err2)

				c.JSON(http.StatusUnauthorized, gin.H{
					"code": code,
					"msg":  "redis中无数据",
					"data": data,
				})
				c.Abort()
				return
			}

			json.Unmarshal(data, &cacheToken)

			if token != cacheToken {
				log.Error("异地登陆")

				c.JSON(http.StatusUnauthorized, gin.H{
					"code": code,
					"msg":  "您的账号目前已经异地登陆了，请重新登陆",
					"data": nil,
				})
				c.Abort()
				return
			}
		}

		if code != constants.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  error3.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		// 成功之后 调用next 说明通过了这个中间件 和node 比较类似
		c.Next()
	}
}
