package model

type PostTag struct {
	Id         int `xorm:"pk autoincr"`
	PostId     uint
	TagId      uint
	CreateTime int `xorm:"created"`
	UpdateTime int `xorm:"updated"`
}

func (p *PostTag) Set() (int64, error) {
	n, err := db.Insert(&p)
	return n, err
}

func (p *PostTag) Get() (bool, error) {
	res, err := db.Get(p)
	return res, err
}

func SetTags(postId int, tags []int) error {
	for _, t := range tags {
		p := PostTag{
			PostId: uint(postId),
			TagId:  uint(t),
		}
		_, err := db.Insert(&p)
		if err != nil {
			return err
		}
	}
	return nil
}
