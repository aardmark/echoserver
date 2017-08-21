package users

import (
	"github.com/aardmark/echoserver/db"
	"github.com/labstack/echo"
)

// Get gets all the users
func Get(c echo.Context) error {
	ds := c.Get("ds").(*db.DataStore)
	result, err := ds.GetUsers()
	if err != nil {
		return err
	}
	return c.JSON(200, result)
}
