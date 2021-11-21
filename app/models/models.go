package models

import "github.com/golang-jwt/jwt"

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
