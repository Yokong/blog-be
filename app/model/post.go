package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Content   []byte `gorm:"type:blob"`
	Title     string `gorm:"size:20"`
	Introduce string `gorm:"size:1024"`
	CoverUrl  string `gorm:"size:120"`
}

func (p *Post) SetWithTags(tags []int) error {
	tx := db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	// 保存文章
	if err := tx.Create(p).Error; err != nil {
		tx.Rollback()
		return err
	}

	//// 保存标签
	if err := SetTags(int(p.ID), tags); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (p *Post) Get(id int) (bool, error) {
	if err := db.Where("id = ?", id).Find(&p).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (p *Post) UpdateById(id int) error {
	p.ID = uint(id)
	err := db.Model(&p).Updates(p).Error
	return err
}

func GetPostWithTitle(title string) *Post {
	var p Post
	if err := db.Where("title = ?", title).Find(&p).Error; err != nil {
		fmt.Println(err)
		return nil
	}
	return &p
}

func GetPostDescList() ([]Post, error) {
	var postDescList []Post
	err := db.Select([]string{"id", "title", "cover_url", "introduce", "created_at", "updated_at"}).Order("created_at", true).Find(&postDescList).Error
	return postDescList, err
}
