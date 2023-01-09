package core_auth

import (
	"fmt"

	log "github.com/BM-laoli/go-gin-example/src/core/log"
	"github.com/BM-laoli/go-gin-example/src/models"
)

type UserType struct {
	Username string `json:"name" valid:"Required; MaxSize(50)"`
	Password string `json:"password" valid:"Required; MaxSize(50)"`
}

// 增加用户
func (a *UserType) AddAuthUser() error {
	article := map[string]interface{}{
		"name":     a.Username,
		"password": a.Password,
	}

	if err := models.AddAuth(article); err != nil {
		return err
	}

	return nil
}

// 看看是否上正确的可以 通过bcript签名通过的用户
func (a *UserType) VerifyUser() (bool, int, error) {
	article := map[string]interface{}{
		"name":     a.Username,
		"password": a.Password,
	}
	log.Info("错误", fmt.Sprintf("--->", article))
	findUser, err := models.GetUserById(article["name"].(string), article["password"].(string))
	log.Info("获取到的用户是", fmt.Sprintf("findUser--->", findUser))
	return models.PasswordVerify(article["password"].(string), findUser.Password), findUser.ID, err
}
