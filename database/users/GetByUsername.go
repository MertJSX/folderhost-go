package users

import (
	"database/sql"

	"github.com/MertJSX/folder-host-go/database"
	"github.com/MertJSX/folder-host-go/types"
)

func GetUserByUsername(username string) (types.Account, error) {
	const query = `
		SELECT
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
			extract_permission,
			copy_permission,
			read_recovery_permission,
			use_recovery_permission
		FROM users
		WHERE username = ?
	`

	row := database.DB.QueryRow(query, username)

	u := types.Account{}

	err := row.Scan(
		&u.Username,
		&u.Password,
		&u.Email,
		&u.Permissions.ReadDirectories,
		&u.Permissions.ReadFiles,
		&u.Permissions.Create,
		&u.Permissions.Change,
		&u.Permissions.Delete,
		&u.Permissions.Move,
		&u.Permissions.DownloadFiles,
		&u.Permissions.UploadFiles,
		&u.Permissions.Rename,
		&u.Permissions.Extract,
		&u.Permissions.Copy,
		&u.Permissions.ReadRecovery,
		&u.Permissions.UseRecovery,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.Account{}, sql.ErrNoRows
		}
		return types.Account{}, err
	}

	return u, nil
}
