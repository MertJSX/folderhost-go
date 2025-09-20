package database

import (
	"fmt"
	"log"

	"github.com/MertJSX/folder-host-go/types"
)

func CheckIfUsernameExists(username string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)"
	err := DB.QueryRow(query, username).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking if username exists: %v", err)
	}
	return exists, nil
}

func CreateUser(user types.Account) error {
	tx, err := DB.Begin()
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

func CreateRecoveryRecord(record types.RecoveryRecord) error {
	tx, err := DB.Begin()
	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("Begin transaction error: %w", err)
	}

	stmt, err := tx.Prepare(`
		INSERT INTO recovery(
			username,
			oldLocation,
			binLocation,
			isDirectory,
			sizeDisplay,
			sizeBytes
		) VALUES(?, ?, ?, ?, ?, ?)
	`)

	if err != nil {
		return fmt.Errorf("error creating db stmt")
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		record.Username,
		record.OldLocation,
		record.BinLocation,
		record.IsDirectory,
		record.SizeDisplay,
		record.SizeBytes,
	)

	if err != nil {
		return fmt.Errorf("error executing db stmt")
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("error commiting db changes")
	}

	return nil
}
