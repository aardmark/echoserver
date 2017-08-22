package db

import (
	"github.com/aardmark/echoserver/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// ErrNotFound indicates no documents were returned
var ErrNotFound = mgo.ErrNotFound

// DataStore abstracts the database
type DataStore struct {
	session *mgo.Session
}

// NewDataStore creates a new data store
func NewDataStore() *DataStore {
	return &DataStore{masterStore.session.Copy()}
}

var masterStore *DataStore

func init() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	masterStore = &DataStore{session}
}

// Close closes the session
func (ds *DataStore) Close() {
	ds.session.Close()
}

// GetUsers gets all the users
func (ds *DataStore) GetUsers() (*[]model.User, error) {
	collection := ds.session.DB("invoicer").C("users")

	var results []model.User
	err := collection.Find(nil).All(&results)
	return &results, err
}

// GetUserByEmail gets a user by email address
func (ds *DataStore) GetUserByEmail(email string) (*model.User, error) {
	collection := ds.session.DB("invoicer").C("users")

	var results model.User
	err := collection.Find(bson.M{"email": email}).One(&results)
	return &results, err
}

// CreateUser created a new user
func (ds *DataStore) CreateUser(user *model.User) error {
	err := ds.session.DB("invoicer").C("users").Insert(user)
	return err
}
