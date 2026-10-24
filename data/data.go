package data

import (
	"io"
	"time"

	"niceygy.net/nawst/panicHandler"
)

// var LoadedEntries []WebViewEntry

func CreateNewEntry(body io.ReadCloser) {
	new := WebViewEntry{}
	new.FromJSONBody(body)
	// LoadedEntries = append(LoadedEntries, new)
	// log.Printf("Got new entry:\nDomain=%s\nPath=%s\nCountry=%s\nHasVisitedBefore=%t", new.Domain, new.Path, new.Country, new.HasVisitedBefore)
	db, dbCloser := OpenDatabase()

	defer dbCloser()

	var monday time.Time = time.Now()

	for monday.Weekday().String() != "Monday" {
		monday = monday.Add((time.Hour * -24))
	}

	res := db.QueryRow("SELECT COUNT(*) FROM pages WHERE domain = ? AND path = ?;", new.Domain, new.Path)
	var rows int
	res.Scan(&rows)

	var hasRow bool = rows > 0

	// fmt.Printf("HasRow=%t", hasRow)
	var err error

	if hasRow {
		_, err = db.Exec("UPDATE pages SET visits = visits + 1 WHERE domain = ? AND path = ? AND weekStarting = ?", new.Domain, new.Path, monday.Format(time.DateOnly))
	} else {
		_, err = db.Exec("INSERT INTO pages (domain, path, weekStarting, visits, viewers) VALUES (?, ?, ?, ?, ?)",
			new.Domain, new.Path, monday.Format(time.DateOnly), 1, 1)
	}
	panicHandler.PanicIfErr("CreateNewEntry", err)
}
