package app

import (
	"github.com/BM-laoli/go-gin-example/pkg/logging"
	"github.com/astaxie/beego/validation"
)

// 阿
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Info(err.Key, err.Message)
	}

	return
}
