package users

import (
	"database/sql"

	"github.com/MertJSX/folder-host-go/database"
)

func GetUsername(id int) (string, error) {
	const query = `
		SELECT
			username
		FROM users
		WHERE id = ?
	`

	row := database.DB.QueryRow(query, id)

	var username string

	err := row.Scan(
		&username,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", sql.ErrNoRows
		}
		return "", err
	}

	return username, nil
}
