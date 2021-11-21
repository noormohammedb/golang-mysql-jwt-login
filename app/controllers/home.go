package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/noormohammedb/golang-mysql-jwt-login/app/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	log.Println("Home")
	cookie, err := r.Cookie("token")
	if err != nil {
		// if err == http.ErrNoCookie {
		// 	log.Println("no token cookie found")
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	w.Write([]byte("unauthorized"))
		// 	return
		// }
		log.Println("no cookie found")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("unauthorized"))
		return
	}

	tokenString := cookie.Value

	authClaim := &models.TokenClaims{}

	jwtObj, err := jwt.ParseWithClaims(tokenString, authClaim, func(jwtToken *jwt.Token) (interface{}, error) {
		return jwtTokenKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			log.Println("jwt signature invalid")
		}
		log.Println("jwt parse error")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("unauthorized"))
		return
	}

	if !jwtObj.Valid {
		log.Print("jwt token invalid")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("unauthorized"))
	}

	w.Write([]byte(fmt.Sprintf("hello, %s", authClaim.Username)))
}
