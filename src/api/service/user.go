package service

import (
	"Goprojects/better-exchange-back/src/api/dao"
	"Goprojects/better-exchange-back/src/api/model"
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