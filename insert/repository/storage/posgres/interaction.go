package posgres

import (
	"database/sql"
	"fmt"
	"new/insert/repository/storage"
	"new/pkg/e"

	_ "github.com/lib/pq"
)

type PosgreSQL struct {
	DB *sql.DB
}

func Connection(dataSourceName *string) (db *sql.DB, err error) {
	db, err = sql.Open("postgres", *dataSourceName)
	if err != nil {
		return nil, e.Err("cen't connect posgres", err)
	}
	fmt.Println("postgres connect readi")
	return db, nil
}

func (db *PosgreSQL) Add(events []storage.Event) (err error) {
	return
}

func (db *PosgreSQL) Update(event *storage.Event) (err error) {
	return
}

func (db *PosgreSQL) All() (events []storage.Event, err error) {
	return
}

func (db *PosgreSQL) Last() (event *storage.Event, err error) {
	return
}
