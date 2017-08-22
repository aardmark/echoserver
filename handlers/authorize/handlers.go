package authorize

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// Get authorizes the client and returns a JWT
func Get(c echo.Context) error {
	token := jwt.New(jwt.SigningMethodHS256)
	c.Logger().Debug(c.Get("user"))
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
