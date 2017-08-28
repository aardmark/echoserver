package users

import (
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"golang.org/x/crypto/bcrypt"

	"github.com/aardmark/echoserver/db"
	"github.com/aardmark/echoserver/model"
	"github.com/labstack/echo"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	return string(bytes), err
}

// Get gets all the users
func Get(c echo.Context) error {
	ds := c.Get("ds").(*db.DataStore)
	result, err := ds.GetUsers()
	if err != nil {
		return err
	}
	return c.JSON(200, result)
}

// Post creates a new user
func Post(c echo.Context) error {
	var err error
	user := &model.User{ID: bson.NewObjectId()}
	uwp := &model.UserWithPassword{User: *user, Password: ""}

	if err = c.Bind(uwp); err != nil {
		return err
	}

	if err = c.Validate(uwp); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}

	if uwp.Password, err = hashPassword(uwp.Password); err != nil {
		return err
	}

	ds := c.Get("ds").(*db.DataStore)
	if err := ds.CreateUser(uwp); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, uwp.User)
}
