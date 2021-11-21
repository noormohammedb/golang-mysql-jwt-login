package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/noormohammedb/golang-mysql-jwt-login/app/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	log.Println("Login")
	var cred models.Credentials
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

	expTimeToken := time.Now().Add(time.Minute * 5)      // expire after 5 minutes of creation
	expTimeRefresh := time.Now().Add(time.Hour * 24 * 5) // expire after 5 days of creation

	tokenClaim := &models.TokenClaims{
		Username: cred.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTimeToken.Unix(),
		},
	}

	refreshClaim := &models.RefreshClaims{
		Username:       cred.Username,
		IsRefreshToken: true,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTimeRefresh.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, tokenClaim)
	refToken := jwt.NewWithClaims(jwt.SigningMethodHS512, refreshClaim)

	// log.Println(token)
	tokenString, err := token.SignedString(jwtTokenKey)
	if err != nil {
		log.Printf("\ntoken signing error %v", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error from login")
		return
	}

	refTokenString, err := refToken.SignedString(jwtRefKey)
	if err != nil {
		log.Printf("\nRefresh token signing error %v", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error from login")
		return
	}

	// log.Println("signed jwt token : ", tokenString)
	// log.Println("signed refresh token : ", refTokenString)

	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expTimeToken})

	http.SetCookie(w,
		&http.Cookie{
			Name:    "refresh",
			Value:   refTokenString,
			Expires: expTimeRefresh})

	// log.Printf("parsed json %#v \n", cred)
	log.Print("login success, jwt and ref tokens created\n")
	fmt.Fprintf(w, "Login Success")
}
