package main

import (
	"github.com/labstack/echo/v5"
	middlewares "github.com/labstack/echo/v5/middleware"
	"niceygy.net/nawst/server"
	"niceygy.net/nawst/web"
)

func main() {
	e := echo.New()

	e.Use(middlewares.CORS("*", "https://*.niceygy.net"))

	e.Any("/ui/*", web.HandleWeb)
	e.Any("/script/*", server.HandleServer)

	e.Start("0.0.0.0:8089")
}
