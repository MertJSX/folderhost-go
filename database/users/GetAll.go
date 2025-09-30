package users

import (
	"github.com/MertJSX/folder-host-go/database"
	"github.com/MertJSX/folder-host-go/types"
)

func GetAll() ([]types.Account, error) {
	const query = `
		SELECT
			username,
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
			use_recovery_permission,
			read_users_permission,
			edit_users_permission,
			logs_permission
		FROM users ORDER BY id;`

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []types.Account

	for rows.Next() {
		u := types.Account{}

		err := rows.Scan(
			&u.Username,
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
			&u.Permissions.ReadUsers,
			&u.Permissions.EditUsers,
			&u.Permissions.ReadLogs,
		)
		if err != nil {
			return nil, err
		}

		u.Password = ""

		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if users == nil {
		return []types.Account{}, nil
	}

	return users, nil
}
