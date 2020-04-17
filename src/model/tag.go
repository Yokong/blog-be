package model

type Tag struct {
	Id         int    `xorm:"pk autoincr"`
	Name       string `gorm:"size:15"`
	CreateTime int    `xorm:"created"`
	UpdateTime int    `xorm:"updated"`
}
