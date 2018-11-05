package service

import (
	"better-exchange-backend/src/model"
	"grpc-rest-api/src/api/dao"
)

var (
	userDAO = dao.NewConnection("localhost:27017", "shaitan")
)

func Register(user model.User) error {
	return userDAO.Insert(user)
}

func Login(data model.LoginData) (bool, error) {
	valid := false
	user, err := userDAO.FindByEmail(data.Email)
	if user.Password == data.Password {
		valid = true
	}
	return valid, err
}

func ChangeProfileData(user model.User)  {
	userDAO.Update(user)
}