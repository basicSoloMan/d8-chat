package main

import (
	"github.com/basicSoloMan/d8-chat/src/api"
	"github.com/basicSoloMan/d8-chat/src/config"
	"log"
	"net/http"
)

func main() {
	router := api.Mount()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	srv := &http.Server{
		Addr:         cfg.Server.Address,
		Handler:      router,
		WriteTimeout: cfg.Server.WriteTimeout,
		ReadTimeout:  cfg.Server.ReadTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	log.Printf("Server listening on %s", cfg.Server.Address)

	log.Fatalln(srv.ListenAndServe())
}
