package core_upload

import (
	"net/http"

	core_constants "github.com/BM-laoli/go-gin-example/src/core/constants"
	core_error "github.com/BM-laoli/go-gin-example/src/core/error"
	logging "github.com/BM-laoli/go-gin-example/src/core/log"
	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {
	code := core_constants.SUCCESS
	data := make(map[string]string)

	file, image, err := c.Request.FormFile("image") // 读取form-data中的数据

	if err != nil {
		logging.Warn(err)
		code = core_constants.ERROR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  core_error.GetMsg(code),
			"data": data,
		})
	}

	if image == nil {
		code = core_constants.INVALID_PARAMS
	} else {
		// 合法图片将会进行 下一步的处理
		imageName := GetImageName(image.Filename)
		fullPath := GetImageFullPath()
		savePath := GetImagePath()

		src := fullPath + imageName
		if !CheckImageExt(imageName) || !CheckImageSize(file) {
			code = core_constants.ERROR_UPLOAD_CHECK_IMAGE_FORMAT
		} else {
			err := CheckImage(fullPath)
			if err != nil {
				logging.Warn(err)
				code = core_constants.ERROR_UPLOAD_CHECK_IMAGE_FAIL
				// gin能够通过这个方法的调用 进行保存  这几个方法调用成功就存上了
			} else if err := c.SaveUploadedFile(image, src); err != nil {
				logging.Warn(err)
				code = core_constants.ERROR_UPLOAD_SAVE_IMAGE_FAIL
			} else {
				data["image_url"] = GetImageFullUrl(imageName)
				data["image_save_url"] = savePath + imageName
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  core_error.GetMsg(code),
		"data": data,
	})
}
