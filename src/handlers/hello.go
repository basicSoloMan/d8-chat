package handlers

import (
	"log"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status": "Hello World",
	}

	if err := writeJSON(w, http.StatusOK, data); err != nil {
		log.Println(err)
	}
}
