package api

import (
	"log"
	"net/http"

	"github.com/BM-laoli/go-gin-example/models"
	"github.com/BM-laoli/go-gin-example/pkg/e"
	"github.com/BM-laoli/go-gin-example/pkg/logging"
	util "github.com/BM-laoli/go-gin-example/utils"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token

				code = e.SUCCESS
			}

		} else {

			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("你大爷的你开始失败了")
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
