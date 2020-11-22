package main

import (
	"blog-be/app/config"
	"blog-be/app/model"
	"blog-be/app/router"
	"log"
	"net/http"
	"time"
)

func main() {
	if err := config.InitConfig(); err != nil {
		panic(err)
	}
	if err := model.InitMode(); err != nil {
		panic(err)
	}
	r := router.InitRouter()

	c := config.GetConfig()

	s := http.Server{
		Addr:              c.ServerAddr,
		Handler:           r,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		IdleTimeout:       15 * time.Second,
	}

	log.Fatal(s.ListenAndServe())
}
