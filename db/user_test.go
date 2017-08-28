package db

import (
	"os"
	"testing"

	"gopkg.in/mgo.v2/bson"

	"github.com/aardmark/echoserver/model"
)

func TestMain(m *testing.M) {
	DBName = "invoicer_test"
	ds := NewDataStore()
	ds.session.DB(DBName).DropDatabase()
	coll := ds.session.DB(DBName).C("users")
	coll.EnsureIndexKey("email")
	coll.Insert(fred)

	results := m.Run()

	ds.session.DB(DBName).DropDatabase()
	os.Exit(results)
}

func TestGetUserByEmail(t *testing.T) {
	ds := NewDataStore()
	f, err := ds.GetUserByEmail("fred@bedrock.gov")
	if err != nil {
		t.Error(err)
		return
	}
	if f.Email != fred.Email {
		t.Error("wrong record returned")
	}
}

func TestCreateUser(t *testing.T) {
	ds := NewDataStore()
	uwp := getUser("foo@bar.org")
	if err := ds.CreateUser(uwp); err != nil {
		t.Error(err)
		return
	}
}

func TestSetUserPassword(t *testing.T) {
	ds := NewDataStore()
	user := getUser("sdf@doo.com")
	ds.CreateUser(user)

	if err := ds.SetUserPassword("sdf@doo.com", "changed"); err != nil {
		t.Error(err)
		return
	}
	// also covers GetUserPassword
	pwd, err := ds.GetUserPassword("sdf@doo.com")
	if err != nil {
		t.Error(err)
	}
	if pwd != "changed" {
		t.Error("password didn't change")
	}
}


// setup stuff and helpers

var fred = model.User{
	FirstName: "Fred",
	LastName:  "Flintstone",
	Email:     "fred@bedrock.gov",
	IsAdmin:   true,
}

var fredWithPassword = model.UserWithPassword{
	User:     fred,
	Password: "$2a$10$fjLgV3E0xjS54.AdDNEX4.ZeYfD1oqhzkJrVuNi82YVPGOa9gLGtu",
}

func getUser(email string) *model.UserWithPassword {
	user := &model.User{ID: bson.NewObjectId()}
	uwp := &model.UserWithPassword{User: *user, Password: "pass"}
	uwp.FirstName = "fname"
	uwp.LastName = "lname"
	uwp.IsAdmin = false
	uwp.Email = email
	return uwp
}
