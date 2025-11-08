package routes

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/MertJSX/folder-host-go/database/logs"
	"github.com/MertJSX/folder-host-go/types"
	"github.com/MertJSX/folder-host-go/utils"
	"github.com/MertJSX/folder-host-go/utils/config"
	"github.com/gofiber/fiber/v2"
)

func CreateCopy(c *fiber.Ctx) error {
	var (
		path       string = c.Query("path")
		parentPath string
		basename   string
		copyPath   string
		extname    string            = ""
		account    types.Account     = c.Locals("account").(types.Account)
		config     *types.ConfigFile = &config.Config
	)

	if !account.Permissions.Copy {
		return c.Status(403).JSON(fiber.Map{"err": "No permission"})
	}

	pathStat, err := os.Stat(fmt.Sprintf("%s%s", config.Folder, path))

	if os.IsNotExist(err) {
		return c.Status(400).JSON(fiber.Map{"err": "The item doesn't exist!"})
	}

	parentPath = utils.GetParentPath(path)
	basename = fmt.Sprintf("%s - Copy", utils.GetPureFileName(path))
	var index int = 0
	if !pathStat.IsDir() {
		extname = filepath.Ext(path)
		copyPath = fmt.Sprintf("%s/%s%s", parentPath, basename, extname)
		if config.StorageLimit != "" {
			fileSize := pathStat.Size()
			remainingFreeSpace, err := utils.GetRemainingFolderSpace()
			if err != nil {
				return c.Status(520).JSON(fiber.Map{"err": "Internal server error!"})
			}

			if fileSize > remainingFreeSpace {
				return c.Status(507).JSON(fiber.Map{"err": "Not enough space!"})
			}
		}

		for utils.IsExistingPath(config.Folder + copyPath) {
			index++
			copyPath = fmt.Sprintf("%s/%s (%d)%s", parentPath, basename, index, extname)
		}

		err := utils.CopyFile(config.Folder+path, config.Folder+copyPath)

		if err != nil {
			return c.Status(520).JSON(fiber.Map{"err": "Internal server error!"})
		}
	} else {
		if config.StorageLimit != "" {
			folderSize, _, err := utils.GetDirectorySize(config.Folder + path)
			if err != nil {
				return c.Status(520).JSON(fiber.Map{"err": "Internal server error!"})
			}
			remainingFreeSpace, err := utils.GetRemainingFolderSpace()
			if err != nil {
				return c.Status(520).JSON(fiber.Map{"err": "Internal server error!"})
			}

			if folderSize > remainingFreeSpace {
				return c.Status(507).JSON(fiber.Map{"err": "Not enough space!"})
			}
		}

		copyPath := fmt.Sprintf("%s/%s", parentPath, basename)

		for utils.IsExistingPath(config.Folder + copyPath) {
			index++
			copyPath = fmt.Sprintf("%s/%s (%d)", parentPath, basename, index)
		}

		if err := utils.CopyDirectory(config.Folder+path, config.Folder+copyPath); err != nil {
			return c.Status(520).JSON(fiber.Map{"err": "Internal server error!"})
		}

		logs.CreateLog(types.AuditLog{
			Username:    c.Locals("account").(types.Account).Username,
			Action:      "Create copy",
			Description: fmt.Sprintf("%s created a copy of %s", c.Locals("account").(types.Account).Username, path),
		})
	}

	return c.Status(200).JSON(fiber.Map{"err": "Copied!"})
}
