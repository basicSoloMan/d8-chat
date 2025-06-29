package main

import (
	"github.com/basicSoloMan/d8-chat/src/api"
	"log"
	"net/http"
	"time"
)

func main() {
	router := api.Mount()

	srv := &http.Server{
		Addr:         "localhost:8080",
		Handler:      router,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server listening on %s", "localhost:8080")

	err := srv.ListenAndServe()
	if err != nil {
		return
	}
}
