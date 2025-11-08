package routes

import (
	"fmt"
	"os"

	"github.com/MertJSX/folder-host-go/types"
	"github.com/MertJSX/folder-host-go/utils"
	"github.com/MertJSX/folder-host-go/utils/config"
	"github.com/gofiber/fiber/v2"
)

func ReadFile(c *fiber.Ctx) error {
	var path string
	var fileName string
	var lastModified string
	var itemStat os.FileInfo
	var err error
	config := &config.Config

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

	content, err := os.ReadFile(fmt.Sprintf("%s%s", config.Folder, path))

	if err != nil {
		fmt.Printf("Error while reading file: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"err": "Error while reading file!"})
	}

	return c.Status(200).JSON(fiber.Map{
		"data":            string(content),
		"res":             "Successfully readed!",
		"title":           fileName,
		"lastModified":    lastModified,
		"writePermission": c.Locals("account").(types.Account).Permissions.Change,
	})
}
