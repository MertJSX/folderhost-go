package routes

import (
	"fmt"
	"os"
	"time"

	"github.com/MertJSX/folder-host-go/types"
	"github.com/MertJSX/folder-host-go/utils"
	"github.com/MertJSX/folder-host-go/utils/cache"
	"github.com/MertJSX/folder-host-go/utils/config"
	"github.com/gofiber/fiber/v2"
)

func GetDownloadLink(c *fiber.Ctx) error {

	if !c.Locals("account").(types.Account).Permissions.DownloadFiles {
		return c.Status(403).JSON(
			fiber.Map{"err": "No permission!"},
		)
	}

	filepath := c.Query("filepath")
	scope := c.Locals("account").(types.Account).Scope
	filepath = fmt.Sprintf("%s%s", config.Config.GetScopedFolder(scope), filepath)

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

	randomID := utils.GenerateUniqueString()

	cache.DownloadLinkCache.Set(randomID, types.DownloadLinkCache{Path: filepath, Username: c.Locals("account").(types.Account).Username}, 1*time.Minute)

	return c.Status(200).JSON(
		fiber.Map{"id": randomID},
	)
}
