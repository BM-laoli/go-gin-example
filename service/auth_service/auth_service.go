package auth_service

import (
	"fmt"

	"github.com/BM-laoli/go-gin-example/models"
	"github.com/BM-laoli/go-gin-example/pkg/logging"
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
func (a *UserType) VerifyUser() (bool, error) {
	article := map[string]interface{}{
		"name":     a.Username,
		"password": a.Password,
	}
	logging.Info("错误", fmt.Sprintf("--->", article))
	findUser, err := models.GetUserById(article["name"].(string), article["password"].(string))
	logging.Info("获取到的用户是", fmt.Sprintf("findUser--->", findUser))
	return models.PasswordVerify(article["password"].(string), findUser.Password), err
}
