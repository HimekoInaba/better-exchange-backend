package service

import (
	"grpc-rest-api/src/api/dao"
	"grpc-rest-api/src/api/model"
)

var (
	userDAO = dao.UserDAO{}
)

func Register(user model.User) error {
	userDAO.Insert(user)
}

func ChangeProfileData(user model.User)  {
	userDAO.Update(user)
}