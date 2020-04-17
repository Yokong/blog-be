package model

type Post struct {
	Id         int    `xorm:"pk autoincr"`
	Title      string `xorm:"varchar(20)"`
	Content    []byte
	CoverUrl   string `xorm:"varchar(120)"`
	CreateTime int    `xorm:"created"`
	UpdateTime int    `xorm:"updated"`
}

func (p *Post) Insert() (int64, error) {
	n, err := db.Insert(p)
	return n, err
}

func (p *Post) Select() (bool, error) {
	res, err := db.Get(p)
	return res, err
}
