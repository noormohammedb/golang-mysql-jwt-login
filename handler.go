package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtTokenKey = []byte("jwt_secret_key")

var jwtRefKey = []byte("jwt_refresh_secret_key")

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type RefreshClaims struct {
	Username       string `json:"username"`
	IsRefreshToken bool   `json:"isRefresh"`
	jwt.StandardClaims
}

func home(w http.ResponseWriter, r *http.Request) {
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

	authClaim := &TokenClaims{}

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

	expTimeToken := time.Now().Add(time.Minute * 5)      // expire after 5 minutes of creation
	expTimeRefresh := time.Now().Add(time.Hour * 24 * 5) // expire after 5 days of creation

	tokenClaim := &TokenClaims{
		Username: cred.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTimeToken.Unix(),
		},
	}

	refreshClaim := &RefreshClaims{
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

func refresh(w http.ResponseWriter, r *http.Request) {
	log.Println("Refresh")

	// cookie, err := r.Cookie("refresh")
	// if err != nil {
	// 	// if err == http.ErrNoCookie {
	// 	// 	log.Println("no token cookie found")
	// 	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	// 	w.Write([]byte("unauthorized"))
	// 	// 	return
	// 	// }
	// 	log.Println("no cookie found")
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	w.Write([]byte("unauthorized"))
	// 	return
	// }

	// tokenString := cookie.Value

	// authClaim := &TokenClaims{}

	// jwtObj, err := jwt.ParseWithClaims(tokenString, authClaim, func(jwtToken *jwt.Token) (interface{}, error) {
	// 	return jwtTokenKey, nil
	// })
	// if err != nil {
	// 	if err == jwt.ErrSignatureInvalid {
	// 		log.Println("jwt signature invalid")
	// 	}
	// 	log.Println("jwt parse error")
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	w.Write([]byte("unauthorized"))
	// 	return
	// }

	// if !jwtObj.Valid {
	// 	log.Print("jwt token invalid")
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	w.Write([]byte("unauthorized"))
	// }

	fmt.Fprintf(w, "hello from refresh")

}
