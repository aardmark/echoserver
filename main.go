package main

import (
	"net/http"

	"github.com/aardmark/echoserver/db"
	"github.com/aardmark/echoserver/handlers/users"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func basicAuthenticator(username, password string, c echo.Context) (bool, error) {
	if username == "fred" && password == "flintstone" {
		c.Set("user", username)
		return true, nil
	}
	return false, nil
}

var trailingSlashConfig = middleware.TrailingSlashConfig{
	RedirectCode: http.StatusMovedPermanently,
}

func main() {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlashWithConfig(trailingSlashConfig))

	e.Logger.SetLevel(log.DEBUG)
	e.Use(db.DataStoreMiddleware())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// e.Use(middleware.JWTWithConfig(jwtConfig))
	// e.Use(middleware.BasicAuthWithConfig(basicAuthConfig))

	e.GET("/authorize", authorize, middleware.BasicAuth(basicAuthenticator))
	e.GET("/accounts", restricted, middleware.JWT([]byte("secret")))
	e.GET("/users", users.Get, middleware.JWT([]byte("secret")))
	e.Logger.Fatal(e.Start(":8080"))
}
