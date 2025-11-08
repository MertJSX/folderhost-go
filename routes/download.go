package routes

import (
	"fmt"
	"os"

	"github.com/MertJSX/folder-host-go/database/logs"
	"github.com/MertJSX/folder-host-go/types"
	"github.com/MertJSX/folder-host-go/utils/config"
	"github.com/gofiber/fiber/v2"
)

func Download(c *fiber.Ctx) error {

	if !c.Locals("account").(types.Account).Permissions.DownloadFiles {
		return c.Status(403).JSON(
			fiber.Map{"err": "No permission!"},
		)
	}

	config := &config.Config
	path := c.Query("filepath")
	filepath := fmt.Sprintf("%s%s", config.Folder, path)

	fileinfo, err := os.Stat(filepath)

	// Validation to avoid errors
	if os.IsNotExist(err) {
		return c.JSON(
			fiber.Map{"err": "Wrong filepath!"},
		)
	} else if fileinfo.IsDir() {
		return c.JSON(
			fiber.Map{"err": "You can't download a directory!"},
		)
	}

	logs.CreateLog(types.AuditLog{
		Username:    c.Locals("account").(types.Account).Username,
		Action:      "Download",
		Description: fmt.Sprintf("%s downloaded a %s file.", c.Locals("account").(types.Account).Username, path),
	})

	return c.Status(200).Download(filepath)
}
