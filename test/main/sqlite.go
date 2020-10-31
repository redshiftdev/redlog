package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"github.com/redshiftdev/redlog"
)

type DB struct {
	conn *sql.DB
	stmt *sql.Stmt
}

func NewDatabaseDriver(path string) (*DB, error) {
	conn, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("could not open sqlite connection: %v", err)
	}
	return &DB{conn: conn}, nil
}

func (db *DB) Start() error {
	_, err := db.conn.Exec(
		"CREATE TABLE IF NOT EXISTS redlog (" +
			"id INTEGER PRIMARY KEY AUTOINCREMENT," +
			"level TEXT NOT NULL," +
			"message TEXT NOT NULL" +
			")",
	)
	if err != nil {
		return fmt.Errorf("could not migrate log table: %v", err)
	}
	stmt, err := db.conn.Prepare("INSERT INTO redlog(`level`, `message`) values(?, ?)")
	if err != nil {
		return fmt.Errorf("could not prepare write statement: %v", err)
	}
	db.stmt = stmt
	return nil
}

func (db *DB) Write(level redlog.LogType, message string) {
	_, err := db.stmt.Exec(level.String(), message)
	if err != nil {
		fmt.Println("failed to write database log:", err)
	}
}

func (db *DB) Close() error {
	err := db.stmt.Close()
	if err != nil {
		return fmt.Errorf("could not close statement: %v (%v)", err, db.conn.Close())
	}
	err = db.conn.Close()
	if err != nil {
		return fmt.Errorf("could not close database: %v", err)
	}
	return nil
}
