package main

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/aardmark/echoserver/db"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var trailingSlashConfig = middleware.TrailingSlashConfig{
	RedirectCode: http.StatusMovedPermanently,
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
	user, err := ds.GetUserByEmail(email)
	if err != nil {
		if err == db.ErrNotFound {
			return false, nil
		}
		c.Logger().Error(err)
		return false, err
	}
	if checkPasswordHash(password, user.Password) {
		c.Set("authenticated_user", user)
		return true, nil
	}
	return false, nil
}
