package users

import (
	"fmt"

	"github.com/MertJSX/folder-host-go/database"
)

func CheckIfUsernameExists(username string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)"
	err := database.DB.QueryRow(query, username).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking if username exists: %v", err)
	}
	return exists, nil
}
