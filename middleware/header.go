package middleware

import (
	"github.com/labstack/echo"
)

func Header(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Content-Type", "application/json")
		return next(c)
	}
}
