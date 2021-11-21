package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/noormohammedb/golang-mysql-jwt-login/app/models"
)

func Refresh(w http.ResponseWriter, r *http.Request) {
	log.Println("Refresh")

	refreshCookie, err := r.Cookie("refresh")

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

	refreshTokenString := refreshCookie.Value

	jwtRefreshClaim := &models.RefreshClaims{}

	jwtRefTokenObj, err := jwt.ParseWithClaims(refreshTokenString, jwtRefreshClaim, func(jwtToken *jwt.Token) (interface{}, error) {
		return jwtRefKey, nil
	})

	log.Print("token data ", jwtRefreshClaim)

	// log.Print("jwt token error")
	// log.Print(jwtTokenObj.Claims.Valid())

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			log.Println("jwt signature invalid")
		}
		log.Println("jwt parse error")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("unauthorized"))
		return
	}

	// if jwtRefTokenObj.Valid {
	// 	log.Print("jwt token valid")
	// 	log.Print(jwtRefTokenObj.Claims.Valid())
	// 	w.Write([]byte("valid"))
	// }

	if jwtRefTokenObj.Valid && jwtRefreshClaim.IsRefreshToken {
		log.Print("refresh token is valid and creating new jwt")
		expTimeToken := time.Now().Add(time.Minute * 5) // expire after 5 minutes of creation

		newTokenClaim := &models.TokenClaims{
			Username: jwtRefreshClaim.Username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expTimeToken.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS512, newTokenClaim)

		// log.Printf("%#v", jwtRefreshClaim)
		tokenString, err := token.SignedString(jwtTokenKey)
		if err != nil {
			log.Printf("\ntoken signing error %v", err)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "error from refresh")
			return
		}
		http.SetCookie(w,
			&http.Cookie{
				Name:    "token",
				Value:   tokenString,
				Expires: expTimeToken})
		fmt.Fprintf(w, "Login Success")
		return
	}

	fmt.Fprintf(w, "Token Refresh Successfull")

}
