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

/*
func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	exp := claims["exp"].(float64)
	admin := claims["admin"].(bool)
	ret := fmt.Sprintf("%v %v %v", email, exp, admin)

	return c.String(200, ret)
}
*/

// Post creates a new user
func Post(c echo.Context) error {
	user := &model.User{ID: bson.NewObjectId()}

	err := c.Bind(user)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}
	err = c.Validate(user)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}

	user.Password, err = hashPassword(user.Password)
	if err != nil {
		return err
	}

	ds := c.Get("ds").(*db.DataStore)
	if err = ds.CreateUser(user); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, user)
}
