package server

import (
	"embed"

	"github.com/labstack/echo/v5"
	"niceygy.net/nawst/data"
	"niceygy.net/nawst/utils"
)

//go:embed widget.js
var JSWidgetContent embed.FS
var JSWidgetText string = ""

func HandleServer(c *echo.Context) error {
	// log.Println(c.Request().RequestURI)
	switch c.Request().RequestURI {
	case "/script/ingest":
		data.CreateNewEntry(c.Request().Body)
		return c.String(200, "OK")

	case "/script/js":
		if len(JSWidgetText) < 1 {
			var renderData map[string]any = map[string]any{
				"scriptIngestURL": "http://localhost:8089/script/ingest",
			}
			JSWidgetText = utils.RenderTemplate(JSWidgetContent, "widget.js", renderData)
		}
		c.Response().WriteHeader(200)
		c.Response().Header().Add("Content-Type", "application/javascript")
		c.Response().Write([]byte(JSWidgetText))
		return nil

	default:
		return c.NoContent(404)
	}
}
