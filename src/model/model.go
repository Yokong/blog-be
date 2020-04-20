package model

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var db *xorm.Engine

func InitMode() {
	var err error
	addr := "root:password@(localhost:3306)/blog?charset=utf8&parseTime=True&loc=Local"
	db, err = xorm.NewEngine("mysql", addr)
	if err != nil {
		panic(err)
	}

	db.ShowSQL(true)

	syncTable(new(Post))
	syncTable(new(Tag))
	syncTable(new(PostTag))
}

func syncTable(tb interface{}) {
	if err := db.Sync2(tb); err != nil {
		fmt.Println(err)
	}
}
