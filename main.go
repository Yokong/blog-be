package main

import (
	"blog-be/src/config"
	"blog-be/src/model"
	"blog-be/src/router"
	"log"
	"net/http"
	"time"
)

func main() {
	model.InitMode()
	r := router.InitRouter()

	s := http.Server{
		Addr:              config.ServerAddr,
		Handler:           r,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		IdleTimeout:       15 * time.Second,
	}

	log.Fatal(s.ListenAndServe())
}
