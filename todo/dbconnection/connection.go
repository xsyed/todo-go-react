package dbconnection

import (
	"todolist/utils"

	mgo "gopkg.in/mgo.v2"
)

type db *mgo.Database

var (
	instance db
)

const (
	HostName       string = "localhost:27017"
	DbName         string = "todo"
	CollectionName string = "todolist"
	Port           string = ":9000"
)

func Connect() db {

	if instance == nil {
		sess, err := mgo.Dial(HostName)
		utils.CheckErr(err)
		sess.SetMode(mgo.Monotonic, true)
		instance = sess.DB(DbName)
	}

	return instance
}
