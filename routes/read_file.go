package routes

import (
	"fmt"
	"os"

	"github.com/MertJSX/folder-host-go/types"
	"github.com/MertJSX/folder-host-go/utils"
	"github.com/MertJSX/folder-host-go/utils/cache"
	"github.com/gofiber/fiber/v2"
)

// Known problems:
//   - Having problems with big files. Needs size limitation.
//     If someone tries to read a file that has 1 GB size, your PC will crash xd.
//   - Real-time checking filesize.
//   - We must ignore big editor-change requests from WebSocket connections.
func ReadFile(c *fiber.Ctx) error {
	var path string
	var fileName string
	var lastModified string
	var itemStat os.FileInfo
	var err error
	config := &utils.Config

	if !c.Locals("account").(types.Account).Permissions.ReadFiles {
		return c.Status(403).JSON(fiber.Map{"err": "No permission!"})
	}

	if c.Query("filepath") == "" {
		return c.Status(400).JSON(fiber.Map{"err": "Bad request!"})
	}

	path = c.Query("filepath")

	if utils.IsNotExistingPath(fmt.Sprintf("%s%s", config.Folder, path)) {
		return c.Status(400).JSON(fiber.Map{"err": "Filepath is not existing!"})
	}

	itemStat, err = os.Stat(fmt.Sprintf("%s%s", config.Folder, path))

	fileName = itemStat.Name()
	lastModified = itemStat.ModTime().GoString()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"err": "Unknown server error!"})
	}

	if itemStat.IsDir() {
		return c.Status(400).JSON(fiber.Map{"err": "Filepath is directory!"})
	}

	if itemStat.Size() > 200*1024 {
		return c.Status(413).JSON(fiber.Map{"err": "File is too large!"})
	}

	if remainingSize, _ := utils.GetRemainingFolderSpace(); remainingSize < 200*1024 {
		return c.Status(413).JSON(fiber.Map{"err": "Not enough storage space to edit! Try to close unused CodeEditor windows. Each code editor window guarantees itself 200 KB of space."})
	}

	var content string
	var ok bool

	if content, ok = cache.FileContentCache.Get(fmt.Sprintf("%s%s", config.Folder, path)); !ok {
		fileContent, err := os.ReadFile(fmt.Sprintf("%s%s", config.Folder, path))
		content = string(fileContent)

		if err != nil {
			return c.Status(500).JSON(fiber.Map{"err": "Error while reading file!"})
		}
	}

	return c.Status(200).JSON(fiber.Map{
		"data":            string(content),
		"res":             "Successfully readed!",
		"title":           fileName,
		"lastModified":    lastModified,
		"writePermission": c.Locals("account").(types.Account).Permissions.Change,
	})
}
