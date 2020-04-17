package model

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var db *xorm.Engine

func init() {
	var err error
	addr := "root:wuyukang.@/blog?charset=utf8&parseTime=True&loc=Local"
	db, err = xorm.NewEngine("mysql", addr)
	if err != nil {
		panic(err)
	}

	db.ShowSQL(true)
}

func InitMode() {
	syncTable(new(Post))
	syncTable(new(Tag))
	syncTable(new(PostTag))
}

func syncTable(tb interface{}) {
	if err := db.Sync2(tb); err != nil {
		fmt.Println(err)
	}
}
