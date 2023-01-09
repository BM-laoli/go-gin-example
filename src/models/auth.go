package models

import (
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) bool {
	// 坚持的方式不应该这样
	// 它应该 只要redis中有这个登录的key就可以证明 用户当前的登录了的
	var auth Auth
	db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	if auth.ID > 0 {
		return true
	}

	return false
}

func GetUserById(username, password string) (Auth, error) {
	var (
		auth Auth
		err  error
	)

	err = db.Where(Auth{Username: username}).First(&auth).Error
	if err != nil {
		return auth, err
	}
	return auth, nil
}

// user注册
func AddAuth(data map[string]interface{}) error {

	pst, err := PasswordHash(data["password"].(string))

	if err != nil {
		return err
	}

	user := &Auth{
		Username: data["name"].(string),
		Password: pst,
	}

	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// 密码加密: 传入密码  ==> 加密后的密码
func PasswordHash(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), 4)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

// 密码验证: 用户发送过来的密码 + 数据库中查出来的密码===>  bool
func PasswordVerify(pwd string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}
