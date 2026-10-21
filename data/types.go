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
	Domain           string
	Path             string
	Fingerprint      string
	DateTime         time.Time
	Country          string
	HasVisitedBefore bool
}

func (w *WebViewEntry) ToString() string {
	dateTimeString := w.DateTime.Format(time.RFC1123)
	domainString := base64.StdEncoding.EncodeToString([]byte(w.Domain))
	pathString := base64.StdEncoding.EncodeToString([]byte(w.Path))
	newVisitorString := ""
	if w.HasVisitedBefore {
		newVisitorString = "T"
	} else {
		newVisitorString = "F"
	}
	return fmt.Sprintf("%s,%s,%s,%s,%s\n", domainString, pathString, dateTimeString, newVisitorString, w.Country)
}

func (w *WebViewEntry) FromString(input string) {
	var inputParts []string = strings.Split(input, ",")
	_domain, err := base64.StdEncoding.DecodeString(inputParts[0])
	panicHandler.PanicIfErr("WebViewEntry.FromString", err)
	w.Domain = string(_domain)

	_path, err := base64.RawStdEncoding.DecodeString(inputParts[1])
	panicHandler.PanicIfErr("WebViewEntry.FromString", err)
	w.Path = string(_path)

	w.DateTime, err = time.Parse(time.RFC1123, inputParts[2])
	w.HasVisitedBefore = inputParts[3] == "T"
	w.Country = inputParts[4]
}

func (w *WebViewEntry) FromJSONBody(body io.ReadCloser) {
	jsonMap := make(map[string]any)
	err := json.NewDecoder(body).Decode(&jsonMap)
	godump.Dump(jsonMap, err)

	w.Country = GetCountryFromTZ(jsonMap["tz"].(string))
	w.DateTime = time.Now()
	w.Domain = jsonMap["dm"].(string)
	w.Path = jsonMap["pt"].(string)
	w.HasVisitedBefore = jsonMap["vb"].(bool)
	w.Fingerprint = jsonMap["fp"].(string)
}
