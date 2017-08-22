package main

import (
	"net/http"

	"github.com/aardmark/echoserver/db"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var trailingSlashConfig = middleware.TrailingSlashConfig{
	RedirectCode: http.StatusMovedPermanently,
}

// DataStoreMiddleware creates middleware to
//  attach a new connection to the request
func DataStoreMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ds := db.NewDataStore()
			defer func() {
				c.Logger().Debug("closing session")
				ds.Close()
			}()
			c.Set("ds", ds)
			return next(c)
		}
	}
}

func basicAuthenticator(username, password string, c echo.Context) (bool, error) {
	if username == "fred" && password == "flintstone" {
		c.Set("user", username)
		return true, nil
	}
	return false, nil
}
