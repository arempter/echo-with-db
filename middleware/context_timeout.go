package middleware

import (
	"context"
	"github.com/labstack/echo/v4"
	"time"
)

func ContextTimeout(timeout time.Duration) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := context.Background()
			ctx, cancel := context.WithTimeout(ctx, timeout)
			defer cancel()

			request := c.Request()
			c.SetRequest(request.WithContext(ctx))

			return next(c)
		}
	}
}
