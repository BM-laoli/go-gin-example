package req

import (
	"net/http"

	error2 "github.com/BM-laoli/go-gin-example/src/core/constants"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// BindAndValid binds and validates data
func BindAndValidForQuery(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form) // 这个是从 ? query 上取值
	if err != nil {
		return http.StatusBadRequest, error2.INVALID_PARAMS
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, error2.ERROR
	}
	if !check {
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, error2.INVALID_PARAMS
	}

	return http.StatusOK, error2.SUCCESS
}

// BindAndValid binds and validates data 从json body 中取值
func BindAndValidForJSON(c *gin.Context, json interface{}) (int, int) {
	err := c.BindJSON(json) // 这里就从body 中读取 json 数据
	if err != nil {
		return http.StatusBadRequest, error2.INVALID_PARAMS
	}

	valid := validation.Validation{}
	check, err := valid.Valid(json)
	if err != nil {
		return http.StatusInternalServerError, error2.ERROR
	}

	if !check {
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, error2.INVALID_PARAMS
	}

	return http.StatusOK, error2.SUCCESS
}
