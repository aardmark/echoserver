package main

import (
	"github.com/aardmark/echoserver/handlers/session"
	"github.com/aardmark/echoserver/handlers/users"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	validator "gopkg.in/go-playground/validator.v9"
)

func main() {
	e := echo.New()
	e.Validator = &customValidator{validator: validator.New()}
	e.Pre(middleware.RemoveTrailingSlashWithConfig(trailingSlashConfig))

	e.Logger.SetLevel(log.DEBUG)
	e.Use(DataStoreMiddleware())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/session", session.Get, middleware.BasicAuth(basicAuthenticator))
	e.GET("/users", users.Get, middleware.JWT([]byte("secret")))
	e.POST("/users", users.Post, middleware.JWT([]byte("secret")))
	e.Logger.Fatal(e.Start(":8080"))
}
