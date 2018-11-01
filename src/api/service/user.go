package service

import (
	"grpc-rest-api/src/api/dao"
	"grpc-rest-api/src/api/model"
)

var (
	userDAO = dao.NewConnection("localhost:27017", "shaitan")
)

func Register(user model.User) error {
	return userDAO.Insert(user)
}

func ChangeProfileData(user model.User)  {
	userDAO.Update(user)
}