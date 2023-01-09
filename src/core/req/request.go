package req

import (
	log "github.com/BM-laoli/go-gin-example/src/core/log"
	"github.com/astaxie/beego/validation"
)

// 阿
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		log.Info(err.Key, err.Message)
	}

	return
}
