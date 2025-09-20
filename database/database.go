package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/MertJSX/folder-host-go/types"
	"github.com/MertJSX/folder-host-go/utils"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitializeDatabase() {
	var err error
	var firstTime bool = false
	if utils.IsNotExistingPath("./database.db") {
		firstTime = true
	}
	DB, err = sql.Open("sqlite3", "./database.db")

	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	if firstTime {
		CreateUsersTable()
		CreateLogsTable()
		CreateRecoveryTable()
		err = CreateUser(types.Account{
			Name:     "admin",
			Password: "admin",
			Permissions: types.AccountPermissions{
				ReadDirectories: true,
				ReadFiles:       true,
				Create:          true,
				Change:          true,
				Delete:          true,
				Move:            true,
				DownloadFiles:   true,
				UploadFiles:     true,
				Rename:          true,
				Unzip:           true,
				Copy:            true,
			},
		})

		if err != nil {
			fmt.Println("Error creating Admin account.")
		}
	}

	fmt.Println("Database connection established successfully!")
}
