package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./steam_comments.db")
	if err != nil {
		fmt.Println(err)
		CreateDatabase()
	}
	return db
}

func CreateDatabase() {
	db := OpenDatabase()
	defer db.Close()

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS messages (
		id INTEGER PRIMARY KEY,
		message_id TEXT NOT NULL
	);
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		fmt.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func InsertMessage(db *sql.DB, message Message) {
	sqlStmt := `
	INSERT INTO messages (message_id) VALUES (?)
	`
	_, err := db.Exec(sqlStmt, message.message_id)
	if err != nil {
		fmt.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func FindMessage(db *sql.DB, message_id string) bool {
	sqlStmt := `
	SELECT EXISTS(SELECT 1 FROM messages WHERE message_id=?)
	`
	var exists bool
	err := db.QueryRow(sqlStmt, message_id).Scan(&exists)
	if err != nil {
		fmt.Printf("%q: %s\n", err, sqlStmt)
		return false
	}
	return exists
}
