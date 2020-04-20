package model

type Post struct {
	Id         int    `xorm:"pk autoincr" json:"id"`
	Title      string `xorm:"varchar(20)" json:"title"`
	Content    []byte `json:"content"`
	Desc       string `xorm:"varchar(1024)" json:"desc"`
	CoverUrl   string `xorm:"varchar(120)" json:"coverUrl"`
	CreateTime int    `xorm:"created" json:"createTime"`
	UpdateTime int    `xorm:"updated" json:"updateTime"`
}

func (p *Post) Set() (int64, error) {
	n, err := db.Insert(p)
	return n, err
}

func (p *Post) Get() (bool, error) {
	res, err := db.Get(p)
	return res, err
}

func GetPostDescList() ([]Post, error) {
	var postDescList []Post
	err := db.Cols("id", "title", "cover_url", "create_time", "update_time").OrderBy("create_time").Find(&postDescList)
	return postDescList, err
}
