package util

import (
	"time"

	"github.com/BM-laoli/go-gin-example/pkg/setting"
	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(setting.AppSetting.JwtSecret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// 生成Token
func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour) // 设置过期时间

	claims := Claims{ // 准备把信息签名 jwt-go 提供
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	// SigningMethodHS256  + Secret加密
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// 解析token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, // 用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
		// 这个是回调 可以多更多的操作
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

	if tokenClaims != nil {
		// tokenClaims.Valid 验证是否过期
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
