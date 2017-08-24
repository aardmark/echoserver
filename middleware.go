package main

import (
	"net/http"

	validator "gopkg.in/go-playground/validator.v9"

	"golang.org/x/crypto/bcrypt"

	"github.com/aardmark/echoserver/db"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var trailingSlashConfig = middleware.TrailingSlashConfig{
	RedirectCode: http.StatusMovedPermanently,
}

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// DataStoreMiddleware creates middleware to
//  attach a new connection to the request
func DataStoreMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ds := db.NewDataStore()
			defer ds.Close()
			c.Set("ds", ds)
			return next(c)
		}
	}
}

func basicAuthenticator(email, password string, c echo.Context) (bool, error) {
	ds := c.Get("ds").(*db.DataStore)
	pwd, err := ds.GetUserPassword(email)
	if err != nil {
		if err == db.ErrNotFound {
			return false, nil
		}
		c.Logger().Error(err)
		return false, err
	}
	if !checkPasswordHash(password, pwd) {
		return false, nil
	}
	user, err := ds.GetUserByEmail(email)
	if err != nil {
		c.Logger().Error(err)
		return false, err
	}
	c.Set("authenticated_user", user)
	return true, nil
}
