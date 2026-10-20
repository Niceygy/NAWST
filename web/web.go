package web

import "github.com/labstack/echo/v5"

func HandleWeb(c *echo.Context) error {
	return c.String(200, "OK")
}
