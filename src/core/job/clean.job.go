package core_job

import (
	logging "github.com/BM-laoli/go-gin-example/src/core/log"
	"github.com/BM-laoli/go-gin-example/src/models"
	"github.com/robfig/cron"
)

func Setup() {
	c := cron.New()

	c.AddFunc("0 0 12 ** ? *", func() { // 每天12 执行一次
		logging.Warn("Run models.CleanAllTag...")
		logging.Warn("Run models.CleanAllArticle...")
		models.CleanAllTag()
		models.CleanAllArticle()
	})

	c.Start()
}
