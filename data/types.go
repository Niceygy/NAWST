package data

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/goforj/godump"
	"niceygy.net/nawst/panicHandler"
)

type WebViewEntry struct {
	domain           string
	path             string
	dateTime         time.Time
	country          string
	hasVisitedBefore bool
}

func (w *WebViewEntry) ToString() string {
	dateTimeString := w.dateTime.Format(time.RFC1123)
	domainString := base64.StdEncoding.EncodeToString([]byte(w.domain))
	pathString := base64.StdEncoding.EncodeToString([]byte(w.path))
	newVisitorString := ""
	if w.hasVisitedBefore {
		newVisitorString = "T"
	} else {
		newVisitorString = "F"
	}
	return fmt.Sprintf("%s,%s,%s,%s,%s\n", domainString, pathString, dateTimeString, newVisitorString, w.country)
}

func (w *WebViewEntry) FromString(input string) {
	var inputParts []string = strings.Split(input, ",")
	_domain, err := base64.StdEncoding.DecodeString(inputParts[0])
	panicHandler.PanicIfErr("WebViewEntry.FromString", err)
	w.domain = string(_domain)

	_path, err := base64.RawStdEncoding.DecodeString(inputParts[1])
	panicHandler.PanicIfErr("WebViewEntry.FromString", err)
	w.path = string(_path)

	w.dateTime, err = time.Parse(time.RFC1123, inputParts[2])
	w.hasVisitedBefore = inputParts[3] == "T"
	w.country = inputParts[4]
}

func (w *WebViewEntry) FromJSONBody(body io.ReadCloser) {
	jsonMap := make(map[string]any)
	err := json.NewDecoder(body).Decode(&jsonMap)
	godump.Dump(jsonMap, err)

	w.country = GetCountryFromTZ(jsonMap["tz"].(string))
	w.dateTime = time.Now()
	w.domain = jsonMap["dm"].(string)
	w.path = jsonMap["pt"].(string)
	w.hasVisitedBefore = jsonMap["vb"].(bool)
}
