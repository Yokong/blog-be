package model

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model
	Name string `gorm:"size:15"`
}

func (t *Tag) Set() error {
	err := db.Create(&t).Error
	return err
}
