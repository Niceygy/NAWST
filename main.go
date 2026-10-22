package main

import (
	"github.com/labstack/echo/v5"
	middlewares "github.com/labstack/echo/v5/middleware"
	"niceygy.net/nawst/data"
	"niceygy.net/nawst/server"
	"niceygy.net/nawst/web"
)

func main() {
	data.MakeTables()
	e := echo.New()

	e.Use(middlewares.CORS("http://localhost:8080", "https://niceyyg.net", "https://*.niceygy.net"))
	e.Use(middlewares.Gzip())

	e.Any("/data/*", web.HandleWeb)
	e.Any("/script/*", server.HandleServer)

	e.Start("0.0.0.0:8089")
}
