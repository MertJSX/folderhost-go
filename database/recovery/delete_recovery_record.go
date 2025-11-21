package recovery

import (
	"fmt"
	"log"

	"github.com/MertJSX/folder-host-go/database"
)

func DeleteRecoveryRecord(id int, scope string) error {
	tx, err := database.DB.Begin()
	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("Begin transaction error: %w", err)
	}

	stmt, err := tx.Prepare(`
		DELETE FROM recovery WHERE id = ? AND oldLocation LIKE ?;
	`)

	if err != nil {
		return fmt.Errorf("error creating db stmt")
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		id,
		scope+"%",
	)

	if err != nil {
		return fmt.Errorf("error executing db stmt")
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("error commiting db changes")
	}

	return nil
}
