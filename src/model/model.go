package model

import (
	"blog-be/src/config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var db *xorm.Engine

func InitMode() {
	var err error
	db, err = xorm.NewEngine("mysql", config.DbAddr)
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
