package model

import "strconv"

const (
	postIdTagNameSql = "select p.post_id, t.name from post_tag p join tag t on t.id = p.tag_id"
)

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

func GetPostIdToTag() map[int][]string {
	res, err := db.Query(postIdTagNameSql)
	if err != nil {
		return nil
	}

	postIdToTag := make(map[int][]string)
	for _, v := range res {
		postIdByte := v["post_id"]
		postId, err := strconv.Atoi(string(postIdByte))
		if err != nil {
			continue
		}

		tags := postIdToTag[postId]
		tags = append(tags, string(v["name"]))
		postIdToTag[postId] = tags
	}

	return postIdToTag
}
