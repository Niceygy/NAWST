package data

import (
	"encoding/json"
	"io"
	"time"

	"github.com/goforj/godump"
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
	godump.Dump(jsonMap, err)

	w.Country = GetCountryFromTZ(jsonMap["tz"].(string))
	w.DateTime = time.Now()
	w.Domain = jsonMap["dm"].(string)
	w.Path = jsonMap["pt"].(string)
	w.VisitedAtAll = jsonMap["vb"].(bool)
	w.VisitedInLastWeek = jsonMap["lw"].(bool)
}
