package controllers

import (
	"database/sql"
	"fmt"

	"github.com/noormohammedb/golang-mysql-jwt-login/config"
)

var jwtTokenKey = []byte("jwt_secret_key")

var jwtRefKey = []byte("jwt_refresh_secret_key")

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

// var refreshTokensArray map[string]string

func getDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.DbCredentials())
	if err != nil {
		panic(err)
	}
	fmt.Println("db connected")
	return db, nil
}
