package mysql

import (
	"database/sql"
	"fmt"
	"new/insert/repository/storage"
	"new/pkg/e"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	DB *sql.DB
}

func Connect(dataSourceName *string) (*sql.DB, error) {
	db, err := sql.Open("mysql", *dataSourceName)
	if err != nil {
		return nil, e.Err("cen't connect mysql", err)
	}

	fmt.Println("mysql connect readi")
	return db, nil
}

func (db *MySQL) Add(events []storage.Event) (err error) {
	return
}

func (db *MySQL) Update(event *storage.Event) (err error) {
	return
}

func (db *MySQL) All() (events []storage.Event, err error) {
	return
}

func (db *MySQL) Last() (event *storage.Event, err error) {
	return
}
