package users

import (
	"fmt"
	"log"

	"github.com/MertJSX/folder-host-go/database"
	"github.com/MertJSX/folder-host-go/types"
)

func CreateUser(user *types.Account) error {
	if exists, _ := CheckIfUsernameExists(user.Username); exists {
		return fmt.Errorf("username already exists")
	}

	tx, err := database.DB.Begin()
	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("Begin transaction error: %w", err)
	}
	if exists, _ := CheckIfUsernameExists(user.Username); exists {
		return fmt.Errorf("username already exists")
	}
	stmt, err := tx.Prepare(`
		INSERT INTO users(
			username,
			password,
			email,
			scope,
			read_directories,
			read_files,
			create_permission,
			change_permission,
			delete_permission,
			move_permission,
			download_permission,
			upload_permission,
			rename_permission,
			extract_permission,
			copy_permission,
			logs_permission,
			read_recovery_permission,
			use_recovery_permission,
			read_users_permission,
			edit_users_permission
		) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)

	if err != nil {
		return fmt.Errorf("error creating db stmt")
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		user.Username,
		user.Password,
		user.Email,
		user.Scope,
		user.Permissions.ReadDirectories,
		user.Permissions.ReadFiles,
		user.Permissions.Create,
		user.Permissions.Change,
		user.Permissions.Delete,
		user.Permissions.Move,
		user.Permissions.DownloadFiles,
		user.Permissions.UploadFiles,
		user.Permissions.Rename,
		user.Permissions.Extract,
		user.Permissions.Copy,
		user.Permissions.ReadLogs,
		user.Permissions.ReadRecovery,
		user.Permissions.UseRecovery,
		user.Permissions.ReadUsers,
		user.Permissions.EditUsers,
	)

	if err != nil {
		return fmt.Errorf("error executing db stmt")
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit error: %w", err)
	}

	return nil
}
