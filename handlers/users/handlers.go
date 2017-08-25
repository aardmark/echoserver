package users

import (
	"encoding/json"
	"io/ioutil"
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
	bodyJSON, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}

	var bodyMap map[string]*json.RawMessage
	err = json.Unmarshal(bodyJSON, &bodyMap)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}

	user := &model.User{ID: bson.NewObjectId()}
	err = json.Unmarshal(bodyJSON, user)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}

	var password string
	err = json.Unmarshal(*bodyMap["password"], &password)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}
	if password == "" {
		c.Logger().Error("Password is required")
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}

	err = c.Validate(user)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}

	password, err = hashPassword(password)
	if err != nil {
		return err
	}

	ds := c.Get("ds").(*db.DataStore)
	if err = ds.CreateUser(user, password); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, user)
}
