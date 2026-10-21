package web

import (
	"strconv"
	"time"

	"github.com/labstack/echo/v5"
	"niceygy.net/nawst/data"
	"niceygy.net/nawst/panicHandler"
)

func HandleWeb(c *echo.Context) error {
	switch c.Request().RequestURI {
	case "/data/byTime":

	}
	data.GetDataForDomain("")
	return c.HTML(200, "")
}

func HandleByTime(c *echo.Context) error {
	hours, err := strconv.Atoi(c.ParamOr("h", "0"))
	panicHandler.PanicIfErr("HandleByTime", err)

	days, err := strconv.Atoi(c.ParamOr("d", "0"))
	panicHandler.PanicIfErr("HandleByTime", err)

	domain := c.ParamOr("domain", "")
	if domain == "" {
		return c.String(404, "Domain not found")
	}

	var oldestEntryTime time.Time = time.Now().Add(-(time.Hour * time.Duration(hours)))
	oldestEntryTime.Add(-((time.Hour * 24) * time.Duration(days)))

	var entries []data.WebViewEntry = data.GetDataForDomain(domain)

	var wantedEntries []data.WebViewEntry = []data.WebViewEntry{}

	for _, e := range entries {
		if e.DateTime.After(oldestEntryTime) {
			wantedEntries = append(wantedEntries, e)
		}
	}
}
