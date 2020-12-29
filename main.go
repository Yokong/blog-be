package main

import (
	"blog-be/app/api"
	"blog-be/app/config"
	"blog-be/app/model"
	"database/sql"
	"log"
)

func main() {
	if err := config.InitConfig(); err != nil {
		panic(err)
	}
	c := config.GetConfig()
	d, err := sql.Open("mysql", c.Db.Addr)
	if err != nil {
		panic(err)
	}

	store := model.NewStore(d)
	s := api.NewServer(store, &c)
	log.Fatal(s.Start())
}
