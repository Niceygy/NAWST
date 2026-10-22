package web

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/labstack/echo/v5"
	"niceygy.net/nawst/data"
	"niceygy.net/nawst/panicHandler"
)

func HandleWeb(c *echo.Context) error {
	// log.Println()
	switch c.Request().URL.Path {
	case "/data/byTime":
		return HandleByTime(c)

	default:
		return c.NoContent(404)
	}
}

func HandleByTime(c *echo.Context) error {
	day, err := strconv.Atoi(c.QueryParamOr("d", "01"))
	panicHandler.PanicIfErr("HandleByTime", err)

	month, err := strconv.Atoi(c.QueryParamOr("m", "01"))
	panicHandler.PanicIfErr("HandleByTime", err)

	year, err := strconv.Atoi(c.QueryParamOr("m", "2026"))
	panicHandler.PanicIfErr("HandleByTime", err)

	domain := c.QueryParamOr("domain", "")
	log.Print(c.QueryParam("domain"))
	// c.QueryParam("domain")
	if domain == "" {
		return c.String(404, "Domain not found")
	}

	var requestedWeekStartingDate string = fmt.Sprintf("%d-%d-%d", day, month, year)

	db, dbCloser := data.OpenDatabase()
	defer dbCloser()

	rows, err := db.Query("SELECT domain, path, vists, viewers FROM pages WHERE domain = ? AND weekStarting = ?", domain, requestedWeekStartingDate)

	panicHandler.PanicIfErr("HandleByTime", err)

	var queryResult []data.WebViewEntry = []data.WebViewEntry{}

	for rows.Next() {
		panicHandler.PanicIfErr("HandleByTime", rows.Err())

		// _row := data.WebViewEntry{}
		// var ignored int
		rows.Scan()
	}

	var result strings.Builder
	fmt.Fprintf(&result, "Domain: %s\n", domain)
	fmt.Fprintf(&result, "Total visits: %d\n", len(queryResult))
	// fmt.Fprintf(&result, "New Users: %d\n", newUsers)
	// fmt.Fprintf(&result, "Total visits: %d\n", totalVisits)

	log.Println(result.String())

	return c.String(200, result.String())
}
