package main

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func authorize(c echo.Context) error {
	token := jwt.New(jwt.SigningMethodHS256)
	fmt.Println(c.Get("user"))
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = "gottogetemail"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	exp := claims["exp"].(float64)
	admin := claims["admin"].(bool)
	ret := fmt.Sprintf("%v %v %v", email, exp, admin)

	return c.String(200, ret)
}
