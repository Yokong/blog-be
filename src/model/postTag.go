package model

type PostTag struct {
	Id         int `xorm:"pk autoincr"`
	PostId     uint
	TagId      uint
	CreateTime int `xorm:"created"`
	UpdateTime int `xorm:"updated"`
}
