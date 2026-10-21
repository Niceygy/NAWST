package data

import (
	"io"
	"log"
	"time"

	"niceygy.net/nawst/panicHandler"
)

// var LoadedEntries []WebViewEntry

func CreateNewEntry(body io.ReadCloser) {
	new := WebViewEntry{}
	new.FromJSONBody(body)
	// LoadedEntries = append(LoadedEntries, new)
	log.Printf("Got new entry:\nDomain=%s\nPath=%s\nCountry=%s\nHasVisitedBefore=%t", new.Domain, new.Path, new.Country, new.HasVisitedBefore)
	db, dbCloser := OpenDatabase()

	defer dbCloser()

	_, err := db.Exec("INSERT INTO visits (domain, path, dateTime, visitedBefore, country) VALUES (?, ?, ?, ?, ?)",
		new.Domain, new.Path, new.DateTime.Format(time.RFC1123), new.HasVisitedBefore, new.Country,
	)

	panicHandler.PanicIfErr("CreateNewEntry", err)
}
