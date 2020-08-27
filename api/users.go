package api

import (
	"echo-with-db/database"
	"echo-with-db/errors"
	"echo-with-db/utils"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserJson struct {
	UserName    string `json:"userName"`
	Description string `json:"description"`
}

func HandlerAddUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		const op errors.Op = "add.user"
		u := new(UserJson)
		db, ok := c.Get("db.service").(database.Database)
		if !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("failed to retrieve db from context"))
		}

		if err := c.Bind(u); err != nil {
			fmt.Println("failed to parse", err)
			return err
		}

		user := database.User{
			UserName:    u.UserName,
			Description: u.Description,
		}

		added := db.AddUser(c.Request().Context(), user)
		if !added {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("failed to add user to db"))
		}
		msg := errors.Message(fmt.Sprintf("Parsing ok, adding db entry, user: %s", u.UserName))
		utils.Log(errors.E(op, msg))

		return c.JSON(http.StatusOK, u)
	}
}
