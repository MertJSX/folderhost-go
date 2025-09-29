package users

import (
	"github.com/MertJSX/folder-host-go/database"
	"github.com/MertJSX/folder-host-go/types"
)

func UpdateUser(user *types.Account) error {
	const query = `
		UPDATE users SET
			password = ?,
			email = ?,
			read_directories = ?,
			read_files = ?,
			create_permission = ?,
			change_permission = ?,
			delete_permission = ?,
			move_permission = ?,
			download_permission = ?,
			upload_permission = ?,
			rename_permission = ?,
			extract_permission = ?,
			copy_permission = ?,
			read_recovery_permission = ?,
			use_recovery_permission = ?
		WHERE username = ?
	`

	_, err := database.DB.Exec(
		query,
		user.Password,
		user.Email,
		user.Permissions.ReadDirectories,
		user.Permissions.ReadFiles,
		user.Permissions.Create,
		user.Permissions.Change,
		user.Permissions.Delete,
		user.Permissions.Move,
		user.Permissions.DownloadFiles,
		user.Permissions.UploadFiles,
		user.Permissions.Rename,
		user.Permissions.Extract,
		user.Permissions.Copy,
		user.Permissions.ReadRecovery,
		user.Permissions.UseRecovery,
		user.Username,
	)

	return err
}
