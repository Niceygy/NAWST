package data

import (
	"database/sql"
	"log"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"niceygy.net/nawst/panicHandler"
)

var DatabaseLock sync.Mutex //not sure if this is needed. We will see
/*CREATE TABLE IF NOT EXISTS visits (
	key INTEGER PRIMARY KEY,
	domain TEXT,
	path TEXT,
	visitedThisWeek BOOL,
	visitedAtAll BOOL
)*/

const TABLE_STATEMENT string = `
CREATE TABLE IF NOT EXISTS pages (
	key INTEGER PRIMARY KEY,
	domain TEXT,
	path TEXT,
	weekStarting TEXT,
	visits INT,
	viewers INT
)
`

func OpenDatabase() (*sql.DB, func()) {
	// DatabaseLock.Lock()
	db, err := sql.Open("sqlite3", "db.sqlite")
	panicHandler.PanicIfErr("OpenDatabase", err)

	dbCloser := func() {
		panicHandler.PanicIfErr("dbCloser", db.Close())
		// DatabaseLock.Unlock()
	}

	return db, dbCloser
}

func MakeTables() {
	db, dbCloser := OpenDatabase()

	defer dbCloser()

	db.Exec(TABLE_STATEMENT)
}

func GetDataForDomain(domain string) []WebViewEntry {
	var result []WebViewEntry = []WebViewEntry{}

	db, dbCloser := OpenDatabase()
	defer dbCloser()

	rows, err := db.Query("SELECT domain, path, visitedThisWeek, visitedAtAll FROM visits;")
	panicHandler.PanicIfErr("GetDataForDomain", err)

	for rows.Next() {
		panicHandler.PanicIfErr("GetDataForDomain", rows.Err())

		var r WebViewEntry
		var _dt string
		if err = rows.Scan(); err == nil {
			r.DateTime, err = time.Parse(time.RFC1123, _dt)
			panicHandler.PanicIfErr("GetDataForDomain", err)
			// godump.Dump(r)
			result = append(result, r)
		} else {
			log.Panic(err)
		}
	}

	return result
}
