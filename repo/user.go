package repo

import (
	"schedule-management/database"
	"schedule-management/model"
)

type UserRepo interface {
	GetListUser() ([]model.User, error)
}

func GetListUsers() (listUser []model.User, err error) {
	err = database.MysqlConn.Find(&listUser).Error
	return
}