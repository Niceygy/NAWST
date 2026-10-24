package main

import (
	"log"
	"os"
	"strings"

	"github.com/labstack/echo/v5"
	middlewares "github.com/labstack/echo/v5/middleware"
	"niceygy.net/nawst/data"
	"niceygy.net/nawst/server"
	"niceygy.net/nawst/web"
)

func LoadAllowedOrigins() []string {
	var origins string = os.Getenv("CORS_ALLOWED")
	if len(origins) < 1 {
		log.Panicf("Expected env CORS_ALLOWED to be a CSV string. Got '%s' instead", origins)
	}
	var originsList []string = strings.Split(origins, ",")
	log.Printf("Loaded allowed origins: %s", origins)
	return originsList
}

func main() {
	data.MakeTables()
	e := echo.New()

	e.Use(middlewares.CORSWithConfig(middlewares.CORSConfig{
		AllowOrigins: LoadAllowedOrigins(),
	}))
	// e.Use(middlewares.CORS("http://localhost:8080", "https://niceygy.net", "https://*.niceygy.net"))
	e.Use(middlewares.Gzip())

	e.Any("/data/*", web.HandleWeb)
	e.Any("/script/*", server.HandleServer)

	e.Start("0.0.0.0:8089")
}
