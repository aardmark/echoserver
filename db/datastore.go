package db

import (
	"github.com/aardmark/echoserver/model"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
)

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

// DataStoreMiddleware creates middleware to attach a new connection
// to the request
func DataStoreMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ds := NewDataStore()
			defer func() {
				ds.session.Close()
			}()
			c.Set("ds", ds)
			return next(c)
		}
	}
}

// GetUsers gets all the users
func (ds *DataStore) GetUsers() ([]model.User, error) {
	defer ds.session.Close()
	collection := ds.session.DB("invoicer").C("users")

	var results []model.User
	err := collection.Find(nil).All(&results)

	return results, err
}
