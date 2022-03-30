package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID        bson.ObjectId `json:"id"`
	FirstName string        `json:"firstname"`
	LastName  string        `json:"lastname"`
	Email     string        `json:"email"`
	Password  string        `json:"password"`
	CreatedAt time.Time     `json:"createAt"`
}
