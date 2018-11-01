package dao

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"grpc-rest-api/src/api/model"
)

type UserDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "user"
)

func NewConnection(server, database string) *UserDAO {
	m := UserDAO{server, database}
	m.Connect()
	return &m
}

func (m *UserDAO) Connect()  {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		panic(err)
	}
	db = session.DB(m.Database)
}

func (m *UserDAO) FindAll() ([]model.User, error) {
	users := []model.User{}
	err := db.C(COLLECTION).Find(bson.M{}).All(&users)

	return users, err
}

func (m *UserDAO) FindById(id string) (model.User, error) {
	var user model.User
	err := db.C(COLLECTION).Find(bson.ObjectIdHex(id)).One(user)
	return user, err
}

func (m *UserDAO) Insert(user model.User) error {
	err := db.C(COLLECTION).Insert(&user)
	if err == nil {
		log.Print("SAKSESFUL")
	}
	return err
}

func (m *UserDAO) Delete(user model.User) error {
	err := db.C(COLLECTION).Remove(&user)
	return err
}

func (m *UserDAO) Update (user model.User) error {
	err := db.C(COLLECTION).Update(user.Id, user)
	return err
}