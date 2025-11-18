package users

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/MertJSX/folder-host-go/database"
	"github.com/MertJSX/folder-host-go/types"
)

func UpdateAdmin(user *types.Account) error {
	const query = `
		UPDATE users SET
			username = ?,
			password = ?,
			email = ?,
			scope = ?,
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
			use_recovery_permission = ?,
			read_users_permission = ?,
			edit_users_permission = ?,
			logs_permission = ?
		WHERE id = 1
	`

	hash := sha256.Sum256([]byte(user.Password))
	hashString := hex.EncodeToString(hash[:])

	_, err := database.DB.Exec(
		query,
		user.Username,
		hashString,
		user.Email,
		user.Scope,
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
		user.Permissions.ReadUsers,
		user.Permissions.EditUsers,
		user.Permissions.ReadLogs,
	)

	return err
}
