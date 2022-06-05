package db

import (
	"database/sql"
	"scheduler/src/config"

	_ "github.com/go-sql-driver/mysql" // Driver
)

func Connect() (*sql.DB, error) {
	db, er := sql.Open("mysql", config.StringConnectionDb)
	if er != nil {
		return nil, er
	}

	if er = db.Ping(); er != nil {
		db.Close()
		return nil, er
	}

	return db, nil
}
