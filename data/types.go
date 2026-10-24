package data

import (
	"encoding/json"
	"io"
	"time"

	"niceygy.net/nawst/panicHandler"
)

type WebViewEntry struct {
	Domain            string
	Path              string
	DateTime          time.Time
	Country           string
	VisitedInLastWeek bool
	VisitedAtAll      bool
}

func (w *WebViewEntry) FromJSONBody(body io.ReadCloser) {
	jsonMap := make(map[string]any)
	err := json.NewDecoder(body).Decode(&jsonMap)
	panicHandler.PanicIfErr("WebViewEntry.FromJSONBody", err)

	if val, ok := jsonMap["tz"].(string); ok {
		w.Country = GetCountryFromTZ(val)
	} else {
		w.Country = "Unknown"
	}

	if val, ok := jsonMap["dm"].(string); ok {
		w.Domain = val
	}

	if val, ok := jsonMap["pt"].(string); ok {
		w.Path = val
	}

	if val, ok := jsonMap["vb"].(bool); ok {
		w.VisitedAtAll = val
	}

	if val, ok := jsonMap["lw"].(bool); ok {
		w.VisitedInLastWeek = val
	}
	w.DateTime = time.Now()

}
