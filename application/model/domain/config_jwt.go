package domain

import "github.com/golang-jwt/jwt"

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var JwtKey = []byte("secret_key")
