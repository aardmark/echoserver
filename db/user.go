package db

import (
	"github.com/aardmark/echoserver/model"
	"gopkg.in/mgo.v2/bson"
)

// GetUsers gets all the users
func (ds *DataStore) GetUsers() (*[]model.User, error) {
	collection := ds.session.DB(DBName).C("users")

	var results []model.User
	err := collection.Find(nil).All(&results)
	return &results, err
}

// GetUserByEmail gets a user by email address
func (ds *DataStore) GetUserByEmail(email string) (*model.User, error) {
	collection := ds.session.DB(DBName).C("users")

	var result model.User
	err := collection.Find(bson.M{"email": email}).One(&result)
	return &result, err
}

// GetUserPassword gets a user's password
func (ds *DataStore) GetUserPassword(email string) (string, error) {
	collection := ds.session.DB(DBName).C("users")
	var result struct {
		Password string `bson:"password"`
	}
	err := collection.Find(bson.M{"email": email}).Select(bson.M{"password": 1}).One(&result)
	return result.Password, err
}

// DeleteUser deletes a user
func (ds *DataStore) DeleteUser(email string) error {
	collection := ds.session.DB(DBName).C("users")
	err := collection.Remove(bson.M{"email": email})
	return err
}

// UpdateUser deletes a user
func (ds *DataStore) UpdateUser(user *model.User) error {
	collection := ds.session.DB(DBName).C("users")
	err := collection.Update(bson.M{"_id": user.ID}, user)
	return err
}

// SetUserPassword sets a user's password
// encryption/hashing should already be done at this point
func (ds *DataStore) SetUserPassword(email, password string) error {
	collection := ds.session.DB(DBName).C("users")
	err := collection.Update(bson.M{"email": email}, bson.M{"$set": bson.M{"password": password}})
	return err
}

// CreateUser creates a new user
func (ds *DataStore) CreateUser(user *model.UserWithPassword) error {
	err := ds.session.DB(DBName).C("users").Insert(user)
	return err
}
