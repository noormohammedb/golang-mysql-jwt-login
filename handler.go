package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("jwt_secret_key")

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func home(w http.ResponseWriter, r *http.Request) {
	log.Println("Home")
	fmt.Println(jwtKey)
	fmt.Fprintf(w, "hello from home")

}

func login(w http.ResponseWriter, r *http.Request) {
	log.Println("Login")
	var cred Credentials
	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil {
		log.Printf("\nrequest body decode error %v", err)
		// log.Printf("%#v",err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error from login")
		return
	}

	userPassFromDb, ok := users[cred.Username]

	if !ok || userPassFromDb != cred.Password {
		log.Print("credentials miss match")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error from login")
		return
	}

	expTime := time.Now().Add(time.Minute * 5)

	newClaim := &Claims{
		Username: cred.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, newClaim)
	log.Println(token)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Printf("\ntoken signing error %v", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error from login")
		return
	}
	log.Println("signed token : ", tokenString)
	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expTime})
	log.Printf("parsed json %#v \n", cred)
	fmt.Fprintf(w, "Login Success")
}

func refresh(w http.ResponseWriter, r *http.Request) {
	log.Println("Refresh")
	fmt.Fprintf(w, "hello from refresh")

}
