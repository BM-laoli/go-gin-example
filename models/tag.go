package models

import "github.com/jinzhu/gorm"

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
// 	// 使用gorm 我们能非常方便的操作mysql 等数据库
// 	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
// 	return
// }

// func GetTagsTotal(maps interface{}) (count int) {
// 	db.Model(&Tag{}).Where(maps).Count(&count)
// 	return
// }

// func ExistTagByName(name string) bool {
// 	var tag Tag
// 	db.Select("id").Where("name = ?", name).First(&tag)
// 	if tag.ID > 0 {
// 		return true
// 	}

// 	return false
// }

// func AddTag(name string, state int, createdBy string) bool {
// 	db.Create(&Tag{
// 		Name:      name,
// 		State:     state,
// 		CreatedBy: createdBy,
// 	})
// 	return true
// }

// func ExistTagByID(id int) bool {
// 	var tag Tag
// 	db.Select("id").Where("id = ?", id).First(&tag)
// 	if tag.ID > 0 {
// 		return true
// 	}

// 	return false

// }

// func DeleteTag(id int) bool {
// 	db.Where("id = ?", id).Delete(&Tag{})
// 	return true
// }

// func EditTag(id int, data interface{}) bool {
// 	db.Model(&Tag{}).Where("id = ?", id).Updates(data)

// 	return true
// }

// // job
// func CleanAllTag() bool {
// 	// 硬删除 Unscoped() 是 gorm约定的
// 	db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Tag{})

// 	return true
// }

// ExistTagByName checks if there is a tag with the same name
func ExistTagByName(name string) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("name = ? AND deleted_on = ? ", name, 0).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if tag.ID > 0 {
		return true, nil
	}

	return false, nil
}

// AddTag Add a Tag
func AddTag(name string, state int, createdBy string) error {
	tag := Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	}
	if err := db.Create(&tag).Error; err != nil {
		return err
	}

	return nil
}

// GetTags gets a list of tags based on paging and constraints
func GetTags(pageNum int, pageSize int, maps interface{}) ([]Tag, error) {
	var (
		tags []Tag
		err  error
	)

	if pageSize > 0 && pageNum > 0 {
		err = db.Where(maps).Find(&tags).Offset(pageNum).Limit(pageSize).Error
	} else {
		err = db.Where(maps).Find(&tags).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return tags, nil
}

// GetTagTotal counts the total number of tags based on the constraint
func GetTagTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Tag{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// ExistTagByID determines whether a Tag exists based on the ID
func ExistTagByID(id int) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}

	return false, nil
}

// DeleteTag delete a tag
func DeleteTag(id int) error {
	if err := db.Where("id = ?", id).Delete(&Tag{}).Error; err != nil {
		return err
	}

	return nil
}

// EditTag modify a single tag
func EditTag(id int, data interface{}) error {
	if err := db.Model(&Tag{}).Where("id = ? AND deleted_on = ? ", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// CleanAllTag clear all tag
func CleanAllTag() (bool, error) {
	if err := db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Tag{}).Error; err != nil {
		return false, err
	}

	return true, nil
}
