package data

import (
	"database/sql"
	"sync"

	_ "github.com/mattn/go-sqlite3"
	"niceygy.net/nawst/panicHandler"
)

var DatabaseLock sync.Mutex //not sure if this is needed. We will see

const TABLE_STATEMENT string = `
CREATE TABLE IF NOT EXISTS visits (
	key INTEGER PRIMARY KEY,
	domain TEXT,
	path TEXT,
	dateTime DATETIME,
	visitedBefore BOOLEAN
)`

func OpenDatabase() (*sql.DB, func()) {
	DatabaseLock.Lock()
	db, err := sql.Open("sqlite3", "db.sqlite")
	panicHandler.PanicIfErr("OpenDatabase", err)

	dbCloser := func() {
		panicHandler.PanicIfErr("dbCloser", db.Close())
		DatabaseLock.Unlock()
	}

	return db, dbCloser
}

func MakeTables() {
	db, dbCloser := OpenDatabase()

	defer dbCloser()

	db.Exec(TABLE_STATEMENT)
}
