package model

import (
	"blog-be/app/config"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func InitMode() error {
	var err error
	c := config.GetConfig()
	db, err = gorm.Open("mysql", c.Db.Addr)
	if err != nil {
		return err
	}

	db.LogMode(true)
	db.SetLogger(log.New(os.Stdout, "\n\r", 0))

	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Tag{})
	db.AutoMigrate(&PostTag{})

	return nil
}
