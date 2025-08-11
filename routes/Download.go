package routes

import (
	"fmt"
	"os"

	"github.com/MertJSX/folder-host-go/types"
	"github.com/MertJSX/folder-host-go/utils"
	"github.com/gofiber/fiber/v2"
)

func Download(c *fiber.Ctx) error {

	if !c.Locals("account").(types.Account).Permissions.DownloadFiles {
		return c.Status(403).JSON(
			fiber.Map{"err": "No permission!"},
		)
	}

	config := utils.GetConfig()
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

	return c.Status(200).Download(filepath)
}
