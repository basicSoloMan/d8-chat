package main

import (
	"log"
	"net/http"
)

func (app *application) helloWorld(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":  "ok",
		"version": version,
	}

	if err := writeJSON(w, http.StatusOK, data); err != nil {
		log.Println(err)
	}
}
