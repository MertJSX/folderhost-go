package logs

import (
	"fmt"

	"github.com/MertJSX/folder-host-go/database"
)

func ResetLogs() error {
	_, err := database.DB.Exec("DELETE FROM logs;")

	if err != nil {
		return fmt.Errorf("error executing db stmt")
	}

	return nil
}
