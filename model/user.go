package model

import "gopkg.in/mgo.v2/bson"

// User is a user of the application
type User struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	FirstName string        `bson:"firstname" json:"firstname" validate:"required"`
	LastName  string        `bson:"lastname" json:"lastname" validate:"required"`
	Email     string        `bson:"email" json:"email" validate:"required,email"`
	IsAdmin   bool          `bson:"isAdmin" json:"isAdmin"`
}
