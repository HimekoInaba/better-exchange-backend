package db

import (
	"better-exchange-backend/src/core/database"
	"better-exchange-backend/src/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

var db *mgo.Database

func (m *database.Connection) FindAll() ([]model.User, error) {
	var users []model.User
	err := db.C(database.COLLECTION).Find(bson.M{}).All(&users)
	return users, err
}

func (m *Connection) FindById(id string) (model.User, error) {
	var user model.User
	err := db.C(database.COLLECTION).Find(bson.ObjectIdHex(id)).One(user)
	return user, err
}

func (m *Connection) Insert(user model.User) error {
	err := db.C(database.COLLECTION).Insert(&user)
	if err == nil {
		log.Print("SAKSESFUL")
	}
	return err
}

func (m *Connection) Delete(user model.User) error {
	err := db.C(database.COLLECTION).Remove(&user)
	return err
}

func (m *Connection) Update (user model.User) error {
	err := db.C(database.COLLECTION).Update(user.Id, user)
	return err
}

func (m *Connection) FindByEmail(email string) (model.User, error)  {
	var user model.User
	err := db.C(COLLECTION).Find(bson.M{"email": email}).One(&user)
	return user, err
}