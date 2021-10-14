package main

import (
	"go-blog-s/internal/routers"
	"net/http"
	"time"
)

func main() {
	router := routers.NewRouter()
	server := http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
