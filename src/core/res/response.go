package res

import (
	error2 "github.com/BM-laoli/go-gin-example/src/core/error"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON  为gin.json设置返格式 类似于OOP中的对类方法的重写
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  error2.GetMsg(errCode),
		Data: data,
	})
	return
}
