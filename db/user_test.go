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
	f, err := ds.GetUserByEmail(fred.Email)
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

func TestDeleteUser(t *testing.T) {
	ds := NewDataStore()
	uwp := getUser("delme@bar.org")
	if err := ds.CreateUser(uwp); err != nil {
		t.Error(err)
		return
	}
	err := ds.DeleteUser("delme@bar.org")
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateUser(t *testing.T) {
	ds := NewDataStore()
	uwp := getUser("update@bar.org")
	if err := ds.CreateUser(uwp); err != nil {
		t.Error(err)
		return
	}
	user := uwp.User
	user.Email = "updated@bar.org"
	user.FirstName = "x"
	user.LastName = "x"
	user.IsAdmin = true
	err := ds.UpdateUser(&user)
	if err != nil {
		t.Error(err)
		return
	}
	u, err := ds.GetUserByEmail("updated@bar.org")
	if err != nil {
		t.Error(err)
		return
	}
	if u.FirstName != "x" || u.LastName != "x" || u.IsAdmin != true {
		t.Error("didn't change")
	}
}

func TestSetAndGetUserPassword(t *testing.T) {
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

func TestGetUsers(t *testing.T) {
	ds := NewDataStore()
	user := getUser("1@doo.com")
	ds.CreateUser(user)
	user = getUser("2@doo.com")
	ds.CreateUser(user)
	users, err := ds.GetUsers()
	if err != nil {
		t.Error(err)
		return
	}
	found := 0
	for _, user := range *users {
		if user.Email == "1@doo.com" || user.Email == "2@doo.com" {
			found++
		}
	}
	if found != 2 {
		t.Error("couldn't find 'em - found ", found)
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
