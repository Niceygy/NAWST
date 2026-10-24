package server

import (
	"embed"
	"net/http"
	"os"

	"github.com/labstack/echo/v5"
	"niceygy.net/nawst/data"
	"niceygy.net/nawst/utils"
)

//go:embed widget.js
var JSWidgetContent embed.FS
var JSWidgetText string = ""

func HandleServer(c *echo.Context) error {
	switch c.Request().URL.Path {
	case "/script/ingest":
		data.CreateNewEntry(c.Request().Body)
		return c.NoContent(http.StatusOK)

	case "/script/js":
		if len(JSWidgetText) < 1 {
			var renderData map[string]any = map[string]any{
				"scriptIngestURL": os.Getenv("PUBLIC_ENDPOINT") + "/script/ingest",
			}
			JSWidgetText = utils.RenderTemplate(JSWidgetContent, "widget.js", renderData)
		}
		c.Response().WriteHeader(200)
		c.Response().Header().Add("Content-Type", "application/javascript")
		c.Response().Write([]byte(JSWidgetText))
		return nil

	default:
		return c.NoContent(http.StatusNotFound)
	}
}
