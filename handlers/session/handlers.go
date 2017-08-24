package session

import (
	"net/http"
	"time"

	"github.com/aardmark/echoserver/model"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// Get authorizes the client and returns a JWT
func Get(c echo.Context) error {
	user := c.Get("authenticated_user").(*model.User)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["firstname"] = user.FirstName
	claims["lastname"] = user.LastName
	claims["email"] = user.Email
	claims["admin"] = user.IsAdmin
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
