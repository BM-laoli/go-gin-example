package api

import (
	"net/http"

	"github.com/BM-laoli/go-gin-example/pkg/app"
	"github.com/BM-laoli/go-gin-example/pkg/e"
	"github.com/BM-laoli/go-gin-example/service/auth_service"
	util "github.com/BM-laoli/go-gin-example/utils"
	"github.com/gin-gonic/gin"
)

type UserType struct {
	Username string `json:"name" valid:"Required; MaxSize(50)"`
	Password string `json:"password" valid:"Required; MaxSize(50)"`
}

// 生产token并返回 类似于登录
func GetAuth(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		user UserType
	)

	httpCode, errCode := app.BindAndValidForJSON(c, &user)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	authService := auth_service.UserType{
		Username: user.Username,
		Password: user.Password,
	}

	// 1.把这个用户名的密码给找出来 从数据库
	// 2. 进行密码核对
	// 3. 若成功就下发token签名

	isValid, _ := authService.VerifyUser()

	if !isValid {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_TAG_FAIL, nil)
		return
	}

	// 开始生产token
	token, err2 := util.GenerateToken(authService.Username, "")
	if err2 != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"token": token,
	})
}

func Register(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		user UserType
	)

	httpCode, errCode := app.BindAndValidForJSON(c, &user)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	authService := auth_service.UserType{
		Username: user.Username,
		Password: user.Password,
	}

	err := authService.AddAuthUser()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
