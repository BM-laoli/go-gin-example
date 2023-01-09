package dto

type AddArticleDto struct {
	TagID         int    `json:"tag_id" valid:"Required;Min(1)"`
	Title         string `json:"title" valid:"Required;MaxSize(100)"`
	Desc          string `json:"desc" valid:"Required;MaxSize(255)"`
	Content       string `json:"content" valid:"Required;MaxSize(65535)"`
	CreatedBy     string `json:"created_by" valid:"Required;MaxSize(100)"`
	CoverImageUrl string `json:"cover_image_url" valid:"Required;MaxSize(255)"`
	State         int    `json:"state" valid:"Range(0,1)"`
}

type EditArticleDto struct {
	ID            int    `json:"id" valid:"Required;Min(1)"`
	TagID         int    `json:"tag_id" valid:"Required;Min(1)"`
	Title         string `json:"title" valid:"Required;MaxSize(100)"`
	Desc          string `json:"desc" valid:"Required;MaxSize(255)"`
	Content       string `json:"content" valid:"Required;MaxSize(65535)"`
	ModifiedBy    string `json:"modified_by" valid:"Required;MaxSize(100)"`
	CoverImageUrl string `json:"cover_image_url" valid:"Required;MaxSize(255)"`
	State         int    `json:"state" valid:"Range(0,1)"`
}

type ArticleQueryDto struct {
	ID    int `form:"id"`
	TagID int `form:"tag_id" valid:"Min(0)"`
	State int `form:"state" valid:"Range(0,1)"`
}
