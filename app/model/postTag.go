package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

const (
	postIdTagNameSql = "select p.post_id, t.name from post_tags p join tags t on t.id = p.tag_id"
)

type PostTag struct {
	gorm.Model
	PostId uint
	TagId  uint
}

func (p *PostTag) Set() error {
	err := db.Create(&p).Error
	return err
}

func (p *PostTag) Get() (bool, error) {
	if err := db.Find(&p).Error; err != nil {
		return false, err
	}
	return true, nil
}

func SetTags(postId int, tags []int) error {
	for _, t := range tags {
		p := PostTag{
			PostId: uint(postId),
			TagId:  uint(t),
		}
		if err := db.Create(&p).Error; err != nil {
			return err
		}
	}
	return nil
}

func GetPostIdToTag() map[int][]string {
	rows, err := db.Raw(postIdTagNameSql).Rows()
	defer rows.Close()
	if err != nil {
		return nil
	}

	postIdToTag := make(map[int][]string)
	for rows.Next() {
		var postId int
		var name string
		if err := rows.Scan(&postId, &name); err != nil {
			fmt.Println(err)
			return nil
		}

		tags := postIdToTag[postId]
		tags = append(tags, name)
		postIdToTag[postId] = tags
	}

	return postIdToTag
}
