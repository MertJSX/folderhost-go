package users

import (
	"github.com/MertJSX/folder-host-go/database"
)

func RemoveUser(id int) error {
	const query = `DELETE FROM users WHERE id = ?;`

	_, err := database.DB.Exec(query, id)
	return err
}
