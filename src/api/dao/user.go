package dao

import (
	"Goprojects/better-exchange-back/src/api/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Connection struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "user"
)

func NewConnection(server, database string) *Connection {
	m := Connection{server, database}
	m.Connect()
	return &m
}

func (m *Connection) Connect()  {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		panic(err)
	}
	db = session.DB(m.Database)
}

func (m *Connection) FindAll() ([]model.User, error) {
	users := []model.User{}
	err := db.C(COLLECTION).Find(bson.M{}).All(&users)

	return users, err
}

func (m *Connection) FindById(id string) (model.User, error) {
	var user model.User
	err := db.C(COLLECTION).Find(bson.ObjectIdHex(id)).One(user)
	return user, err
}

func (m *Connection) Insert(user model.User) error {
	err := db.C(COLLECTION).Insert(&user)
	if err == nil {
		log.Print("SAKSESFUL")
	}
	return err
}

func (m *Connection) Delete(user model.User) error {
	err := db.C(COLLECTION).Remove(&user)
	return err
}

func (m *Connection) Update (user model.User) error {
	err := db.C(COLLECTION).Update(user.Id, user)
	return err
}
func (m *Connection) FindByEmail(email string) (model.User, error)  {
	var user model.User
	err := db.C(COLLECTION).Find(bson.M{"email": email}).One(&user)
	return user, err
}