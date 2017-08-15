package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.JWTWithConfig(
		middleware.JWTConfig{
			Skipper: func(c echo.Context) bool {
				path := c.Path()
				if path == "/login" || path == "/accessible" {
					return true
				}
				return false
			},
			SigningKey: []byte("secret"),
		}))

	e.POST("/login", login)
	e.GET("/accessible", accessible)
	e.GET("/restricted", restricted)
	e.GET("/accounts", restricted)

	e.Logger.Fatal(e.Start(":8080"))
}
