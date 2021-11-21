package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/noormohammedb/golang-mysql-jwt-login/app/controllers"
)

func main() {
	fmt.Println("Golang Mysql JWT Authendication")

	http.HandleFunc("/home", controllers.Home)
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/refresh", controllers.Refresh)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
