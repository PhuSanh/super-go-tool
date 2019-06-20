package handler

import (
	"github.com/labstack/echo/v4"
	"schedule-management/jsonwebtoken"
	"schedule-management/logger"
	"schedule-management/model"
)

type AuthHandler struct {}

func (a AuthHandler) Login(c echo.Context) (err error) {
	var body model.LoginForm
	if err = c.Bind(body); err != nil {
		logger.Error("Login request params wrong")
		return respondToClient(c, 200, "Login request params wrong", nil)
	}
	// TODO: find in model
	authClaims := &model.AuthClaims{
		Id: 1,
		Username: body.Username,
		Role: "admin",
	}
	tokenString, err := jsonwebtoken.Create(authClaims)
	auth := model.AuthResponse{
		Token: tokenString,
		Id: 1,
		Username: body.Username,
		Role: "admin",
	}
	return respondToClient(c, 200, "Get token success", auth)
}

func (a AuthHandler) Check(c echo.Context) (err error) {
	token := c.Request().Header.Get("Authorization")
	claims, _ := jsonwebtoken.Validate(token)
	return respondToClient(c, 200, "Get token success", claims)
}