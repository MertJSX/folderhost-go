package recovery

import (
	"fmt"

	"github.com/MertJSX/folder-host-go/database"
	"github.com/MertJSX/folder-host-go/types"
)

func SearchRecoveryRecords(limit, skip int) ([]types.RecoveryRecord, error) {
	var foundList []types.RecoveryRecord
	rows, err := database.DB.Query(`
		SELECT * FROM recovery ORDER BY created_at DESC LIMIT ? OFFSET ?;
	`, limit, skip)

	if err != nil {
		return nil, fmt.Errorf("error while getting recovery records: %v", err)
	}

	for rows.Next() {
		var record types.RecoveryRecord
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
			return nil, fmt.Errorf("error while getting recovery records: %v", err)
		}

		foundList = append(foundList, record)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error while getting recovery records: %v", err)
	}

	return foundList, nil
}
