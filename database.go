package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./steam_comments.db")
	if err != nil {
		log.Println(err)
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
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func PopulateDatabase(db *sql.DB) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM messages").Scan(&count)
	if err != nil {
		log.Println("Error checking database for messages:", err)
		return
	}

	if count == 0 {
		config := LoadConfig()
		messages := ScrapeSteamGroup(config.GroupName)

		for _, message := range messages {
			InsertMessage(db, message)
		}

		fmt.Println("Inital Database Population Complete")
	}
}

func InsertMessage(db *sql.DB, message Message) {
	sqlStmt := `
	INSERT INTO messages (message_id) VALUES (?)
	`
	_, err := db.Exec(sqlStmt, message.message_id)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
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
		log.Printf("%q: %s\n", err, sqlStmt)
		return false
	}
	return exists
}
