package model

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id bson.ObjectId `bson:"_id" json:"id"`
	Username string `bson:"username" json:"username"`
	Password string `bson:"password"`
	Email string `bosn:"email" json:"email"`
	Avatar []byte `bson:"avatar" json:"avatar"`
	BirthDate string `bson:"birthDate" json:"birth_date"`
}