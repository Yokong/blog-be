package main

import (
	"blog-be/src/router"
	"log"
	"net/http"
	"time"
)

func main() {
	r := router.InitRouter()

	s := http.Server{
		Addr:              ":8989",
		Handler:           r,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		IdleTimeout:       15 * time.Second,
	}

	log.Fatal(s.ListenAndServe())
}
