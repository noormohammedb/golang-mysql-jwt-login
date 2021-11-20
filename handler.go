package main

import (
	"fmt"
	"log"
	"net/http"
)

var jwtKey = []byte("jwt_secret_key")

func home(w http.ResponseWriter, r *http.Request) {
	log.Println("Home")
	fmt.Println(jwtKey)
	fmt.Fprintf(w, "hello from home")

}

func login(w http.ResponseWriter, r *http.Request) {
	log.Println("Login")
	fmt.Fprintf(w, "hello from login")

}

func refresh(w http.ResponseWriter, r *http.Request) {
	log.Println("Refresh")
	fmt.Fprintf(w, "hello from refresh")

}
