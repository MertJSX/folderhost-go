package recovery

import (
	"fmt"
	"log"

	"github.com/MertJSX/folder-host-go/database"
	"github.com/MertJSX/folder-host-go/types"
)

func CreateRecoveryRecord(record types.RecoveryRecord) error {
	tx, err := database.DB.Begin()
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
