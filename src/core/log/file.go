package core_log

import (
	"fmt"
	core_configuration "github.com/BM-laoli/go-gin-example/src/core/configuration"
	core_file "github.com/BM-laoli/go-gin-example/src/core/file"
	"os"
	"time"
)

var (
	LogSavePath = "runtime/logs/"
	LogSaveName = "log"
	LogFileExt  = "log"
	TimeFormat  = "20060102"
)

// 获取文件路径
func getLogFilePath() string {
	return fmt.Sprintf("%s%s", core_configuration.AppSetting.RuntimeRootPath, core_configuration.AppSetting.LogSavePath)
}

// 获取文件名称
func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		core_configuration.AppSetting.LogSaveName,
		time.Now().Format(core_configuration.AppSetting.TimeFormat),
		core_configuration.AppSetting.LogFileExt,
	)
}

// 打开日志文件
func openLogFile(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	src := dir + "/" + filePath
	perm := core_file.CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	err = core_file.IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	f, err := core_file.Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
	}

	return f, nil
}
