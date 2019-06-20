package model

import "github.com/dgrijalva/jwt-go"

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthClaims struct {
	Id 			int64 	`json:"id"`
	Username 	string 	`json:"username"`
	Role 		string 	`json:"role"`
	jwt.StandardClaims
}

type AuthResponse struct {
	Token 		string 	`json:"token"`
	Id 			int64 	`json:"id"`
	Username 	string 	`json:"username"`
	Role 		string 	`json:"role"`
}