package core_auth

import (
	"net/http"
	"strconv"

	error2 "github.com/BM-laoli/go-gin-example/src/core/constants"
	core_redis "github.com/BM-laoli/go-gin-example/src/core/redis"
	"github.com/BM-laoli/go-gin-example/src/core/req"
	"github.com/BM-laoli/go-gin-example/src/core/res"
	"github.com/BM-laoli/go-gin-example/src/dto"

	"github.com/gin-gonic/gin"
)

// 生产token并返回 类似于登录
func GetAuth(c *gin.Context) {
	var (
		appG = res.Gin{C: c}
		user dto.UserDto
	)

	httpCode, errCode := req.BindAndValidForJSON(c, &user)
	if errCode != error2.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	authService := UserType{
		Username: user.Username,
		Password: user.Password,
	}

	// 1.把这个用户名的密码给找出来 从数据库
	// 2. 进行密码核对
	// 3. 若成功就下发token签名

	isValid, uid, _ := authService.VerifyUser()

	if !isValid {
		appG.Response(http.StatusInternalServerError, error2.ERROR_ADD_TAG_FAIL, nil)
		return
	}

	key := "uid_" + strconv.Itoa(uid)

	// 开始生产token
	token, err2 := GenerateToken(authService.Username, "", key)
	if err2 != nil {
		appG.Response(http.StatusInternalServerError, error2.ERROR_ADD_TAG_FAIL, nil)
		return
	}

	// 把user_id 作为key token 做value 存入redis
	core_redis.Set(key, token, 1800) // 注意这个时间需要 >= jwt_token 过期时间

	appG.Response(http.StatusOK, error2.SUCCESS, map[string]interface{}{
		"token": token,
	})
}

func Register(c *gin.Context) {
	var (
		appG = res.Gin{C: c}
		user dto.UserDto
	)

	httpCode, errCode := req.BindAndValidForJSON(c, &user)
	if errCode != error2.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	authService := UserType{
		Username: user.Username,
		Password: user.Password,
	}

	err := authService.AddAuthUser()
	if err != nil {
		appG.Response(http.StatusInternalServerError, error2.ERROR_ADD_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, error2.SUCCESS, nil)
}
