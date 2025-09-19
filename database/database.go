package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitializeDatabase() {
	var err error
	DB, err = sql.Open("sqlite3", "./database/database.db")

	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	CreateUsersTable()
	CreateLogsTable()
	CreateRecoveryTable()

	fmt.Println("Database connection established successfully!")
}

func CreateUsersTable() {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL UNIQUE,
			password TEXT NULL,
			email TEXT NULL,
			read_directories BOOLEAN NOT NULL DEFAULT FALSE,
        	read_files BOOLEAN NOT NULL DEFAULT FALSE,
        	create_permission BOOLEAN NOT NULL DEFAULT FALSE,
        	change_permission BOOLEAN NOT NULL DEFAULT FALSE,
        	delete_permission BOOLEAN NOT NULL DEFAULT FALSE,
        	move_permission BOOLEAN NOT NULL DEFAULT FALSE,
        	download_permission BOOLEAN NOT NULL DEFAULT FALSE,
        	upload_permission BOOLEAN NOT NULL DEFAULT FALSE,
        	rename_permission BOOLEAN NOT NULL DEFAULT FALSE,
        	archive_permission BOOLEAN NOT NULL DEFAULT FALSE,
        	copy_permission BOOLEAN NOT NULL DEFAULT FALSE,
			logs_permission BOOLEAN NOT NULL DEFAULT FALSE,
			recovery_permission BOOLEAN NOT NULL DEFAULT FALSE,
        	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Users table has been created!")
}

func CreateLogsTable() {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS logs (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			user TEXT NOT NULL,
			action TEXT NULL,
        	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Logs table has been created!")

}

func CreateRecoveryTable() {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS recovery (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			user TEXT NOT NULL,
			oldLocation TEXT NULL,
			binLocation TEXT NULL,
			sizeBytes INTEGER NOT NULL DEFAULT 0,
        	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Recovery table has been created!")

}
