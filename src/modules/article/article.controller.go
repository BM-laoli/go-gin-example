package article

import (
	"net/http"

	configuration "github.com/BM-laoli/go-gin-example/src/core/configuration"
	error2 "github.com/BM-laoli/go-gin-example/src/core/constants"
	req2 "github.com/BM-laoli/go-gin-example/src/core/req"
	"github.com/BM-laoli/go-gin-example/src/core/res"
	util "github.com/BM-laoli/go-gin-example/src/core/utils"
	"github.com/BM-laoli/go-gin-example/src/dto"
	"github.com/BM-laoli/go-gin-example/src/modules/tag"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// GetArticle 获取某一个文章
func GetArticle(c *gin.Context) {
	appG := res.Gin{C: c}
	query := dto.ArticleQueryDto{
		ID: com.StrTo(c.Param("id")).MustInt() | 1,
	}

	articleService := Article{ID: query.ID}
	exists, err := articleService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, error2.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, error2.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	article, err := articleService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, error2.ERROR_GET_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, error2.SUCCESS, article)
}

func GetArticles(c *gin.Context) {
	appG := res.Gin{C: c}

	valid := validation.Validation{}
	state := -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state")
	}
	tagId := -1
	if arg := c.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		valid.Min(tagId, 1, "tag_id")
	}

	if valid.HasErrors() {
		req2.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, error2.INVALID_PARAMS, nil)
		return
	}

	articleService := Article{
		TagID:    tagId,
		State:    state,
		PageNum:  util.GetPage(c),
		PageSize: configuration.AppSetting.PageSize,
	}

	total, err := articleService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, error2.ERROR_COUNT_ARTICLE_FAIL, nil)
		return
	}

	articles, err := articleService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, error2.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = articles
	data["total"] = total

	appG.Response(http.StatusOK, error2.SUCCESS, data)
}

func AddArticle(c *gin.Context) {
	var (
		appG = res.Gin{C: c}
		form = dto.AddArticleDto{}
	)

	httpCode, errCode := req2.BindAndValidForJSON(c, &form)
	if errCode != error2.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	tagService := tag.Tag{ID: form.TagID}
	exists, err := tagService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, error2.ERROR_EXIST_TAG_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, error2.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	articleService := Article{
		TagID:         form.TagID,
		Title:         form.Title,
		Desc:          form.Desc,
		Content:       form.Content,
		CoverImageUrl: form.CoverImageUrl,
		State:         form.State,
		CreatedBy:     form.CreatedBy,
	}
	if err := articleService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, error2.ERROR_ADD_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, error2.SUCCESS, nil)
}

func EditArticle(c *gin.Context) {
	var (
		appG     = res.Gin{C: c}
		jsonBody = dto.EditArticleDto{}
	)

	httpCode, errCode := req2.BindAndValidForJSON(c, &jsonBody)
	if errCode != error2.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	articleService := Article{
		ID:            jsonBody.ID,
		TagID:         jsonBody.TagID,
		Title:         jsonBody.Title,
		Desc:          jsonBody.Desc,
		Content:       jsonBody.Content,
		CoverImageUrl: jsonBody.CoverImageUrl,
		ModifiedBy:    jsonBody.ModifiedBy,
		State:         jsonBody.State,
	}
	exists, err := articleService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, error2.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, error2.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	tagService := tag.Tag{ID: jsonBody.TagID}
	exists, err = tagService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, error2.ERROR_EXIST_TAG_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, error2.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	err = articleService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, error2.ERROR_EDIT_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, error2.SUCCESS, nil)
}

func DeleteArticle(c *gin.Context) {
	appG := res.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		req2.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, error2.INVALID_PARAMS, nil)
		return
	}

	articleService := Article{ID: id}
	exists, err := articleService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, error2.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, error2.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	err = articleService.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, error2.ERROR_DELETE_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, error2.SUCCESS, nil)
}
