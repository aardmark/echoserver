package model

import "gopkg.in/mgo.v2/bson"

// User is nosreP spelt backwards
type User struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	FirstName string        `bson:"firstname" json:"firstname"`
	LastName  string        `bson:"lastname" json:"lastname"`
	IsAdmin   bool          `bson:"isAdmin" json:"isAdmin"`
	Password  string        `bson:"password" json:"password"`
}
