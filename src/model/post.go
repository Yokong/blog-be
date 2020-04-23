package model

type Post struct {
	Id         int    `xorm:"pk autoincr" json:"id"`
	Title      string `xorm:"varchar(20)" json:"title" binding:"required"`
	Content    []byte `json:"content"`
	Desc       string `xorm:"varchar(1024)" json:"desc"`
	CoverUrl   string `xorm:"varchar(120)" json:"coverUrl"`
	CreateTime int    `xorm:"created" json:"createTime"`
	UpdateTime int    `xorm:"updated" json:"updateTime"`
}

func (p *Post) SetWithTags(tags []int) error {
	session := db.NewSession()
	err := session.Begin()
	if err != nil {
		return err
	}

	// 保存文章
	_, err = db.Insert(p)
	if err != nil {
		session.Rollback()
		return err
	}

	//// 保存标签
	if err := SetTags(p.Id, tags); err != nil {
		session.Rollback()
		return err
	}

	err = session.Commit()
	return err
}

func (p *Post) Get() (bool, error) {
	res, err := db.Get(p)
	return res, err
}

func GetPostWithTitle(title string) *Post {
	p := Post{
		Title: title,
	}
	ok, err := db.Get(&p)
	if !ok || err != nil {
		return nil
	}
	return &p
}

func GetPostDescList() ([]Post, error) {
	var postDescList []Post
	err := db.Cols("id", "title", "cover_url", "create_time", "update_time").OrderBy("create_time").Find(&postDescList)
	return postDescList, err
}
