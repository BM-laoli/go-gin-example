package dto

type EditTagDto struct {
	ID   int    `form:"id" valid:"Required;Min(1)"`
	Name string `form:"name" valid:"Required;MaxSize(100)"`
	// 后面的定义验证 json 和form 在你把这个结构体 传入gin.bind 的时候会分类绑定 要么form格式，要么json格式
	// 建议统一使用json
	ModifiedBy string `json:"modified_by" valid:"Required;MaxSize(100)"`
	State      int    `form:"state" valid:"Range(0,1)"`
}

type AddTagDto struct {
	Name      string `json:"name" valid:"Required;MaxSize(100)"`
	CreatedBy string `json:"created_by" valid:"Required;MaxSize(100)"`
	State     int    `json:"state" valid:"Range(0,1)"`
}

type TagQueryDto struct {
	Name  string `json:"name"`
	State int    `json:"state"`
}
