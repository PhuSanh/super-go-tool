package jsonwebtoken

import (
	"github.com/dgrijalva/jwt-go"
	"schedule-management/model"
)

func Create(data jwt.Claims) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), data)
	tokenString, err = token.SignedString([]byte("this_is_secret_key"))
	return tokenString, err
}

func Validate(tokenString string) (*model.AuthClaims, string) {
	token, err := jwt.ParseWithClaims(tokenString, &model.AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("this_is_secret_key"), nil
	})

	if err != nil {
		return nil, err.Error()
	}

	if !token.Valid {
		return nil, "Token is not valid."
	}

	claims, _ := token.Claims.(*model.AuthClaims)
	return claims, ""
}