package controllers

import (
	"log"
	"net/http"
)

func AllRoute(w http.ResponseWriter, r *http.Request) {
	log.Println("route not found")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Route Not Found"))
}
