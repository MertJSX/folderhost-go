package recovery

import (
	"fmt"

	"github.com/MertJSX/folder-host-go/database"
	"github.com/MertJSX/folder-host-go/types"
)

func GetRecoveryRecord(id int) (types.RecoveryRecord, error) {
	var record types.RecoveryRecord
	rows, err := database.DB.Query(`
		SELECT * FROM recovery WHERE id = ?;
	`, id)

	if err != nil {
		return record, fmt.Errorf("error while getting recovery records: %v", err)
	}

	for rows.Next() {
		if err := rows.Scan(
			&record.Id,
			&record.Username,
			&record.OldLocation,
			&record.BinLocation,
			&record.IsDirectory,
			&record.SizeDisplay,
			&record.SizeBytes,
			&record.CreatedAt); err != nil {
			fmt.Println(err)
			return record, fmt.Errorf("error while getting recovery records: %v", err)
		}
	}

	if err := rows.Err(); err != nil {
		return record, fmt.Errorf("error while getting recovery records: %v", err)
	}

	return record, nil
}
