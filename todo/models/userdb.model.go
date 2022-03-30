package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type UserModel struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	FirstName string        `bson:"firstname"`
	LastName  string        `bson:"lastname"`
	Email     string        `bson:"email"`
	Password  string        `bson:"password"`
	CreatedAt time.Time     `bson:"createAt"`
}
