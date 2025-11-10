package routes

import (
	"fmt"
	"os"

	"github.com/MertJSX/folder-host-go/database/logs"
	"github.com/MertJSX/folder-host-go/types"
	"github.com/MertJSX/folder-host-go/utils/cache"
	"github.com/gofiber/fiber/v2"
)

func Download(c *fiber.Ctx) error {

	id := c.Query("id")
	downloadLinkCache, ok := cache.DownloadLinkCache.Get(id)

	if !ok {
		return c.Status(400).JSON(fiber.Map{"err": "ID not found"})
	}

	fileinfo, err := os.Stat(downloadLinkCache.Path)

	// Validation to avoid errors
	if os.IsNotExist(err) {
		cache.DownloadLinkCache.Delete(id)
		return c.JSON(
			fiber.Map{"err": "Wrong filepath!"},
		)
	} else if fileinfo.IsDir() {
		cache.DownloadLinkCache.Delete(id)
		return c.JSON(
			fiber.Map{"err": "You can't download a directory!"},
		)
	}

	logs.CreateLog(types.AuditLog{
		Username:    downloadLinkCache.Username,
		Action:      "Download",
		Description: fmt.Sprintf("%s downloaded a %s file.", downloadLinkCache.Username, downloadLinkCache.Path),
	})

	c.Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fileinfo.Name()))
	c.Set("Content-Type", "application/octet-stream")
	c.Set("Content-Length", fmt.Sprintf("%d", fileinfo.Size()))

	return c.SendFile(downloadLinkCache.Path)
}
