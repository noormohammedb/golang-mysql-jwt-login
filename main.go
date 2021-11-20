package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Golang Mysql JWT Authendication")

	http.HandleFunc("/home", home)
	http.HandleFunc("/login", login)
	http.HandleFunc("/refresh", refresh)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
