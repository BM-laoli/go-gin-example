package models

import "github.com/jinzhu/gorm"

//  ------------------------------------------------准备工作定义模型 和callback
type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"` //声明 这个字段是 索引 外键  （⚠️：如果你想关联 那么外键 和 struct 都要有）
	Tag   Tag `json:"tag"`                 //  内嵌的model

	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         int    `json:"state"`
	CoverImageUrl string `json:"cover_image_url"`
}

//  ------------------------------------------------ 正式开始逻辑

// 判断有没有这个文章
func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)

	if article.ID > 0 {
		return true
	}

	return false
}

// 获取LIst总数
func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)

	return
}

// 获取文章List
// func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
// 	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
// 	// 注意这个 Preload
// 	// 它 是一个预加载器 就执行两条SQL，上面的👆的orm翻译过来就是
// 	// SELECT * FROM blog_articles;    SELECT * FROM blog_tag WHERE id IN (1,2,3,4);
// 	// 在gorm中做关联查询主要是两种方式
// 	// 1. gorm 的Join 2. 循环的Related
// 	return
// }
// GetArticles gets a list of articles based on paging constraints
func GetArticles(pageNum int, pageSize int, maps interface{}) ([]*Article, error) {
	var articles []*Article
	err := db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return articles, nil
}

func GetArticle(id int) (*Article, error) {
	var article Article
	err := db.Where("id = ? AND deleted_on = ? ", id, 0).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	err = db.Model(&article).Related(&article.Tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &article, nil
}

// 这里有一个空接口 类似于any
func EditArticle(id int, data interface{}) error {
	if err := db.Model(&Article{}).Where("id = ? AND deleted_on = ? ", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil

}

func AddArticle(data map[string]interface{}) error {
	article := &Article{
		TagID: data["tag_id"].(int),
		// 我们来看看这个语法 ，实际上它想表达的是：
		// 1. V.( I ) 断言 ，I 表示接口interface V表示一个借口值， V.( I ) 这句话的意思是 看看 接口的值 是否为某一个类型
		// 2. 结合上述的理解就是： 看看 类型data中的Tag_id的值 是否是int
		Title:         data["title"].(string),
		Desc:          data["desc"].(string),
		Content:       data["content"].(string),
		CreatedBy:     data["created_by"].(string),
		State:         data["state"].(int),
		CoverImageUrl: data["cover_image_url"].(string),
	}

	if err := db.Create(&article).Error; err != nil {
		return err
	}
	return nil
}

// DeleteArticle delete a single article
func DeleteArticle(id int) error {
	if err := db.Where("id = ?", id).Delete(Article{}).Error; err != nil {
		return err
	}

	return nil
}

// 定时job
func CleanAllArticle() bool {
	db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Article{})

	return true
}
