package logs

import (
	"fmt"

	"github.com/MertJSX/folder-host-go/database"
	"github.com/MertJSX/folder-host-go/types"
)

func SearchLogs(limit, skip int) ([]types.AuditLog, error) {
	var foundList []types.AuditLog
	rows, err := database.DB.Query(`
		SELECT * FROM logs ORDER BY created_at DESC LIMIT ? OFFSET ?;
	`, limit, skip)

	if err != nil {
		return nil, fmt.Errorf("error while getting recovery records: %v", err)
	}

	for rows.Next() {
		var logItem types.AuditLog
		if err := rows.Scan(
			&logItem.ID,
			&logItem.Username,
			&logItem.Action,
			&logItem.Description,
			&logItem.CreatedAt); err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("error while getting recovery records: %v", err)
		}

		foundList = append(foundList, logItem)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error while getting recovery records: %v", err)
	}

	return foundList, nil
}
