package routes

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/MertJSX/folder-host-go/types"
	"github.com/MertJSX/folder-host-go/utils"
	"github.com/gofiber/fiber/v2"
)

func ReadDirectory(c *fiber.Ctx) error {
	var body map[string]interface{}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"err": "Bad request"})
	}

	if !c.Locals("account").(types.Account).Permissions.ReadFiles {
		return c.JSON(
			fiber.Map{"err": "No permission!"},
		)
	}

	path := c.Query("folder")
	fmt.Println(path)
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

	var config types.ConfigFile = utils.GetConfig()
	var dirPath string = fmt.Sprintf("%s%s", config.Folder, path)
	// var fsys fs.FS = os.DirFS("./")
	directoryData, err := os.Stat(dirPath)

	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return c.JSON(
			fiber.Map{"err": "Unknown error!"},
		)
	}

	_, err = os.Stat(dirPath)

	// Validation to avoid errors
	if os.IsNotExist(err) {
		return c.JSON(
			fiber.Map{"err": "Wrong dirpath!"},
		)
	} else if !directoryData.IsDir() {
		return c.JSON(
			fiber.Map{"err": "Dirpath is not directory!"},
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
	var directoryInfo types.DirectoryItem

	if mode() == "Optimized mode" {
		directoryInfo = types.DirectoryItem{
			Name:         folderName,
			ParentPath:   utils.GetParentPath(dirPath),
			Path:         dirPath,
			IsDirectory:  directoryData.IsDir(),
			DateModified: directoryData.ModTime(),
			Size:         "N/A",
			SizeBytes:    directoryData.Size(),
		}
	} else {
		dirSizeBytes, dirSize, _ := utils.GetDirectorySize(fmt.Sprintf("%s%s", config.Folder, path))
		directoryInfo = types.DirectoryItem{
			Name:         folderName,
			ParentPath:   utils.GetParentPath(dirPath),
			Path:         dirPath,
			IsDirectory:  directoryData.IsDir(),
			DateModified: directoryData.ModTime(),
			Size:         dirSize,
			SizeBytes:    dirSizeBytes,
		}
	}

	if config.StorageLimit != "" {
		directoryInfo.StorageLimit = config.StorageLimit
	} else {
		directoryInfo.StorageLimit = "UNLIMITED"
	}

	var data []types.DirectoryItem = utils.GetDirectoryItems(fmt.Sprintf("%s%s", config.Folder, path), mode(), config)

	return c.JSON(
		fiber.Map{
			"data":          data,
			"isEmpty":       len(data) == 0,
			"res":           "Successfully readed!",
			"directoryInfo": directoryInfo,
			"permissions":   c.Locals("account").(types.Account).Permissions,
		},
	)
}
