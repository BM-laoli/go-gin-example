package core_upload

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"

	configuration "github.com/BM-laoli/go-gin-example/src/core/configuration"
	file "github.com/BM-laoli/go-gin-example/src/core/file"
	log2 "github.com/BM-laoli/go-gin-example/src/core/log"
	util "github.com/BM-laoli/go-gin-example/src/core/utils"
)

// 获取图片完整的访问路径 依据名字
func GetImageFullUrl(name string) string {
	return configuration.AppSetting.PrefixUrl + "/" + GetImagePath() + name
}

// 获取图片的名称
func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

// 获取图片的路径
func GetImagePath() string {
	return configuration.AppSetting.ImageSavePath
}

// 获取图片的完整路径
func GetImageFullPath() string {
	return configuration.AppSetting.RuntimeRootPath + GetImagePath()
}

// 检查图片后缀
func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range configuration.AppSetting.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

// 检测图片大小
func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		log2.Warn(err)
		return false
	}

	return size <= configuration.AppSetting.ImageMaxSize
}

// 看看是否是图片
func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
