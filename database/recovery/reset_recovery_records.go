package recovery

import (
	"fmt"

	"github.com/MertJSX/folder-host-go/database"
)

func ResetRecoveryRecords() error {
	_, err := database.DB.Exec("DELETE FROM recovery;")

	if err != nil {
		return fmt.Errorf("error executing db stmt")
	}

	return nil
}
