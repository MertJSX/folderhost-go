package recovery

import (
	"fmt"

	"github.com/MertJSX/folder-host-go/database"
)

func ResetRecoveryRecords(scope string) error {
	_, err := database.DB.Exec("DELETE FROM recovery WHERE oldLocation LIKE ?;", scope+"%")

	if err != nil {
		return fmt.Errorf("error executing db stmt")
	}

	return nil
}
