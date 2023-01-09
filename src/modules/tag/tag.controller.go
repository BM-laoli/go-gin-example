package tag

import (
	configuration "github.com/BM-laoli/go-gin-example/src/core/configuration"
	error2 "github.com/BM-laoli/go-gin-example/src/core/constants"
	req2 "github.com/BM-laoli/go-gin-example/src/core/req"
	"github.com/BM-laoli/go-gin-example/src/core/res"
	util "github.com/BM-laoli/go-gin-example/src/core/utils"
	"github.com/BM-laoli/go-gin-example/src/dto"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetTags(c *gin.Context) {
	appG := res.Gin{C: c}

	query := dto.TagQueryDto{
		Name:  c.Query("name"),
		State: com.StrTo(c.Query("state")).MustInt() | -1,
	}

	tagService := Tag{
		Name:     query.Name,
		State:    query.State,
		PageNum:  util.GetPage(c),
		PageSize: configuration.AppSetting.PageSize,
	}

	// 获取全部
	tags, err := tagService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, error2.ERROR_GET_TAGS_FAIL, nil)
		return
	}

	// 获取总条数
	count, err := tagService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, error2.ERROR_COUNT_TAG_FAIL, nil)
		return
	}

	// 对方法体进行统一格式返回
	appG.Response(http.StatusOK, error2.SUCCESS, map[string]interface{}{
		"lists": tags,
		"total": count,
	})
}

func AddTag(c *gin.Context) {
	var (
		appG     = res.Gin{C: c}
		jsonForm dto.AddTagDto
	)

	// 验证form表单
	httpCode, errCode := req2.BindAndValidForJSON(c, &jsonForm)
	if errCode != error2.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	tagService := Tag{
		Name:      jsonForm.Name,
		CreatedBy: jsonForm.CreatedBy,
		State:     jsonForm.State,
	}

	exists, err := tagService.ExistByName()
	if err != nil {
		appG.Response(http.StatusInternalServerError, error2.ERROR_EXIST_TAG_FAIL, nil)
		return
	}

	if exists {
		appG.Response(http.StatusOK, error2.ERROR_EXIST_TAG, nil)
		return
	}

	err = tagService.Add()
	if err != nil {
		appG.Response(http.StatusInternalServerError, error2.ERROR_ADD_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, error2.SUCCESS, nil)
}

func EditTag(c *gin.Context) {
	var (
		appG     = res.Gin{C: c}
		formJSON dto.EditTagDto
	)

	httpCode, errCode := req2.BindAndValidForJSON(c, &formJSON)
	if errCode != error2.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	tagService := Tag{
		ID:         formJSON.ID,
		Name:       formJSON.Name,
		ModifiedBy: formJSON.ModifiedBy,
		State:      formJSON.State,
	}

	// 看看文章 是否存在 防止 在并发的时候出问题
	exists, err := tagService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, error2.ERROR_EXIST_TAG_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, error2.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	err = tagService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, error2.ERROR_EDIT_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, error2.SUCCESS, nil)
}

func DeleteTag(c *gin.Context) {
	appG := res.Gin{C: c}

	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")
	if valid.HasErrors() {
		req2.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, error2.INVALID_PARAMS, nil)
	}

	tagService := Tag{ID: id}

	exists, err := tagService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, error2.ERROR_EXIST_TAG_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, error2.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	if err := tagService.Delete(); err != nil {
		appG.Response(http.StatusInternalServerError, error2.ERROR_DELETE_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, error2.SUCCESS, nil)
}
