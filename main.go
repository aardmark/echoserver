package main

import (
	"net/http"

	"github.com/aardmark/echoserver/db"
	"github.com/aardmark/echoserver/handlers/users"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

var jwtConfig = middleware.JWTConfig{
	Skipper: func(c echo.Context) bool {
		return c.Path() == "/authorize" || c.Path() == "/users"
	},
	SigningKey: []byte("secret"),
}

var basicAuthConfig = middleware.BasicAuthConfig{
	Skipper: func(c echo.Context) bool {
		return c.Path() != "/authorize"
	},
	Validator: func(username, password string, c echo.Context) (bool, error) {
		if username == "fred" && password == "flintstone" {
			c.Set("user", username)
			return true, nil
		}
		return false, nil
	},
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
	e.Use(middleware.JWTWithConfig(jwtConfig))
	e.Use(middleware.BasicAuthWithConfig(basicAuthConfig))

	e.GET("/authorize", authorize)
	e.GET("/accounts", restricted)
	e.GET("/users", users.Get)
	e.Logger.Fatal(e.Start(":8080"))
}
