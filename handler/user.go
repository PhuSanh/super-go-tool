package handler

import (
	"github.com/labstack/echo/v4"
	"schedule-management/logger"
	"schedule-management/repo"
)

type UserHandler struct {}

// ListUsers godoc
// @Summary List users
// @Description get users
// @Accept  json
// @Produce  json
// @Success 200 {object} handler.ResponseData
// @Failure 404 {object} handler.ResponseData
// @Router /users [get]
func (u UserHandler) GetList(c echo.Context) (err error) {
	users, err := repo.GetListUsers()
	if err != nil {
		return respondToClient(c, 404, "Get users fail", nil)
	}
	logger.Info("Get user success", logger.LogFields{
		Host: "127.0.0.1",
		Source: "127.0.0.1",
		FullMessage: "This is full message get success",
	})
	return respondToClient(c, 200, "Get users success", users)
}