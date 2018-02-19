package main

import (
	"net/http"

	"log"
)

func getHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		w.WriteHeader(http.StatusNotImplemented)
	}
}