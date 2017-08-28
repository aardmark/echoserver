package db

import (
	"gopkg.in/mgo.v2"
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
// DBName specifies the database name
var DBName string

func init() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	masterStore = &DataStore{session}
	DBName = "invoicer"
}

// Close closes the session
func (ds *DataStore) Close() {
	ds.session.Close()
}
