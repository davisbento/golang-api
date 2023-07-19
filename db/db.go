package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	*sql.DB
}

func NewDB() *DB {
	return &DB{}
}

func (db *DB) Connect(path string) (*sql.DB, error) {
	conn, err := sql.Open("sqlite3", path)

	if err != nil {
		return nil, err
	}

	db.DB = conn

	return conn, nil
}
