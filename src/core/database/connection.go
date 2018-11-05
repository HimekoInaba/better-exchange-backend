package database

import "gopkg.in/mgo.v2"

type Connection struct {
	Server   string
	Database string
}

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