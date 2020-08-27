package middleware

import (
	"echo-with-db/database"
	"github.com/labstack/echo/v4"
)

func AttachDatabase(db database.Database) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db.service", db)
			return next(c)
		}
	}
}
