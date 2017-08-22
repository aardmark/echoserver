package main

import (
	"github.com/aardmark/echoserver/handlers/users"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlashWithConfig(trailingSlashConfig))

	e.Logger.SetLevel(log.DEBUG)
	e.Use(DataStoreMiddleware())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/authorize", authorize, middleware.BasicAuth(basicAuthenticator))
	e.GET("/users", users.Get, middleware.JWT([]byte("secret")))
	e.Logger.Fatal(e.Start(":8080"))
}
