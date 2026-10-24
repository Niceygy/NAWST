package web

import (
	"encoding/json/v2"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v5"
	"niceygy.net/nawst/data"
	"niceygy.net/nawst/panicHandler"
)

type ByTimeResult struct {
	Path   string
	Visits int
}

func IsAuthorised(c *echo.Context) bool {
	return c.Request().Header.Get("Authorisation") == os.Getenv("API_PASSWORD")
}

func HandleWeb(c *echo.Context) error {

	if !IsAuthorised(c) {
		return c.String(http.StatusUnauthorized, "Bad Authorisation Password")
	}

	switch c.Request().URL.Path {
	case "/data/lastWeek":
		return HandleLastWeek(c)

	default:
		return c.NoContent(404)
	}
}

func GetLastMonday() time.Time {
	var result time.Time = time.Now()

	for result.Weekday() != time.Monday {
		result = result.Add(time.Hour * -24)
	}

	return result
}

func HandleLastWeek(c *echo.Context) error {
	var lastMon time.Time = GetLastMonday()

	domain := c.QueryParamOr("domain", "")
	if domain == "" {
		return c.String(404, "Domain not found")
	}

	var requestedWeekStartingDate string = lastMon.Format(time.DateOnly)
	var todayDate string = time.Now().Format(time.DateOnly)

	db, dbCloser := data.OpenDatabase()
	defer dbCloser()

	rows, err := db.Query("SELECT path, visits FROM pages WHERE domain = ? AND weekStarting = ? ORDER BY visits DESC;", domain, requestedWeekStartingDate)

	panicHandler.PanicIfErr("HandleByTime", err)

	var queryResult []ByTimeResult = []ByTimeResult{}
	var totalVisits int = 0

	for rows.Next() {
		panicHandler.PanicIfErr("HandleByTime", rows.Err())

		var path string
		var visits int
		rows.Scan(&path, &visits)

		totalVisits += visits
		queryResult = append(queryResult, ByTimeResult{
			Path:   path,
			Visits: visits,
		})
	}

	var jResult map[string]any = map[string]any{
		"domain":      domain,
		"from":        requestedWeekStartingDate,
		"to":          todayDate,
		"entries":     len(queryResult),
		"totalVisits": totalVisits,
		"visits":      queryResult,
	}
	jsonData, err := json.Marshal(jResult, json.DefaultOptionsV2())

	return c.String(200, string(jsonData))
}
