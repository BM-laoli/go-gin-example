package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//  ------------------------------------------------准备工作定义模型 和callback
type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"` //声明 这个字段是 索引 外键  （⚠️：如果你想关联 那么外键 和 strcut都要有）
	Tag   Tag `json:"tag"`                 //  内嵌的model

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
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
func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	// 注意这个 Preload
	// 它 是一个预加载器 就执行两条SQL，上面的👆的orm翻译过来就是
	// SELECT * FROM blog_articles;    SELECT * FROM blog_tag WHERE id IN (1,2,3,4);
	// 在gorm中做关联查询主要是两种方式
	// 1. gorm 的Join 2. 循环的Related
	return
}

func GetArticle(id int) (article Article) {
	// 这个是做关联查询的
	db.Where("id = ?", id).First(&article)   // 先查 文章
	db.Model(&article).Related(&article.Tag) // Related 方式（gorm提供的功能） 再查 它的Tag

	return
}

// 这里有一个空接口 类似于any
func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Updates(data)

	return true
}

func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagID: data["tag_id"].(int),
		// 我们来看看这个语法 ，实际上它想表达的是：
		// 1. V.( I ) 断言 ，I 表示接口interface V表示一个借口值， V.( I ) 这句话的意思是 看看 接口的值 是否为某一个类型
		// 2. 结合上述的理解就是： 看看 类型data中的Tag_id的值 是否是int
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})

	return true
}

func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(Article{})

	return true
}
