package users

import (
	"fmt"
	"log"

	"github.com/MertJSX/folder-host-go/database"
	"github.com/MertJSX/folder-host-go/types"
)

func CreateUser(user types.Account) error {
	tx, err := database.DB.Begin()
	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("Begin transaction error: %w", err)
	}
	if exists, _ := CheckIfUsernameExists(user.Name); exists {
		return fmt.Errorf("username already exists")
	}
	stmt, err := tx.Prepare(`
		INSERT INTO users(
			username,
			password,
			email,
			read_directories,
			read_files,
			create_permission,
			change_permission,
			delete_permission,
			move_permission,
			download_permission,
			upload_permission,
			rename_permission,
			archive_permission,
			copy_permission,
			logs_permission,
			recovery_permission
		) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)

	if err != nil {
		return fmt.Errorf("error creating db stmt")
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		user.Name,
		user.Password,
		"",
		user.Permissions.ReadDirectories,
		user.Permissions.ReadFiles,
		user.Permissions.Create,
		user.Permissions.Change,
		user.Permissions.Delete,
		user.Permissions.Move,
		user.Permissions.DownloadFiles,
		user.Permissions.UploadFiles,
		user.Permissions.Copy,
		false,
		false,
	)

	if err != nil {
		return fmt.Errorf("error executing db stmt")
	}

	return nil
}
