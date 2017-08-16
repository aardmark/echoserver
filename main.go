package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var jwtConfig = middleware.JWTConfig{
	Skipper: func(c echo.Context) bool {
		return c.Path() == "/authorize"
	},
	SigningKey: []byte("secret"),
}

var basicAuthConfig = middleware.BasicAuthConfig{
	Skipper: func(c echo.Context) bool {
		return c.Path() != "/authorize"
	},
	Validator: func(username, password string, c echo.Context) (bool, error) {
		if username == "fred" && password == "flintstonex" {
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

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.JWTWithConfig(jwtConfig))
	e.Use(middleware.BasicAuthWithConfig(basicAuthConfig))

	e.GET("/authorize", authorize)
	e.GET("/accounts", restricted)
	getUserByEmail("fred@bedrock.gov")
	e.Logger.Fatal(e.Start(":8080"))
}
