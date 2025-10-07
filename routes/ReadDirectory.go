package routes

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"syscall"

	"github.com/MertJSX/folder-host-go/types"
	"github.com/MertJSX/folder-host-go/utils"
	"github.com/gofiber/fiber/v2"
)

func ReadDirectory(c *fiber.Ctx) error {
	if !c.Locals("account").(types.Account).Permissions.ReadDirectories {
		return c.Status(403).JSON(
			fiber.Map{"err": "No permission!"},
		)
	}

	path := c.Query("folder")
	mode := func() string {
		switch c.Query("mode") {
		case "Optimized mode":
			return c.Query("mode")
		case "Quality mode":
			return c.Query("mode")
		default:
			return "Balanced mode"
		}
	}

	if utils.LastChar(path) != "/" {
		path += "/"
	}

	config := &utils.Config
	var dirPath string = fmt.Sprintf("%s%s", config.Folder, path)
	directoryData, err := os.Stat(dirPath)

	if os.IsNotExist(err) {
		return c.Status(400).JSON(
			fiber.Map{"err": "Wrong dirpath!"},
		)
	}

	if errors.Is(err, syscall.ENOTDIR) {
		return c.Status(400).JSON(
			fiber.Map{"err": "Dirpath is not a directory!"},
		)
	}

	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{"err": "Unknown error!"},
		)
	}

	trimmedPath := func() string {
		if utils.LastChar(dirPath) == "/" {
			return dirPath[0 : len(dirPath)-1]
		} else {
			return dirPath
		}
	}
	cleanedPath := filepath.Clean(trimmedPath())
	folderName := filepath.Base(cleanedPath)
	dirPath = utils.ReplacePathPrefix(dirPath, fmt.Sprintf("%s/", config.Folder))

	directoryInfo := types.DirectoryItem{
		Name:         folderName,
		ParentPath:   utils.GetParentPath(dirPath),
		Path:         dirPath,
		IsDirectory:  directoryData.IsDir(),
		DateModified: directoryData.ModTime(),
		Size:         "N/A",
		SizeBytes:    directoryData.Size(),
	}

	if config.StorageLimit != "" {
		directoryInfo.StorageLimit = config.StorageLimit
	} else {
		directoryInfo.StorageLimit = "UNLIMITED"
	}

	data, mainDirectorySize := utils.GetDirectoryItems(fmt.Sprintf("%s%s", config.Folder, path), mode())

	if mainDirectorySize != 0 {
		directoryInfo.SizeBytes = mainDirectorySize
		directoryInfo.Size = utils.ConvertBytesToString(mainDirectorySize)
	}

	directoryInfo.Id = -1

	return c.JSON(
		fiber.Map{
			"items":         data,
			"directoryInfo": directoryInfo,
		},
	)
}
