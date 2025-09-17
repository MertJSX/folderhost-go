package routes

import (
	"fmt"
	"os"
	"strconv"

	"github.com/MertJSX/folder-host-go/types"
	"github.com/MertJSX/folder-host-go/utils"
	"github.com/gofiber/fiber/v2"
)

func CreateItem(c *fiber.Ctx) error {
	var (
		itemPath string            = c.Query("path")
		itemName string            = c.Query("itemName")
		account  types.Account     = c.Locals("account").(types.Account)
		config   *types.ConfigFile = &utils.Config
		isFolder bool
	)

	if !account.Permissions.Create {
		return c.Status(403).JSON(
			fiber.Map{"err": "No permission!"},
		)
	}

	if utils.IsExistingPath(fmt.Sprintf("%s%s/%s", config.Folder, itemPath, itemName)) {
		return c.Status(400).JSON(
			fiber.Map{"err": "Item already exists!"},
		)
	}

	isFolder, err := strconv.ParseBool(c.Query("isFolder"))

	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{"err": "isFolder query is not existing or it's not boolean!"},
		)
	}

	if isFolder {
		err = os.Mkdir(fmt.Sprintf("%s%s/%s", config.Folder, itemPath, itemName), 0644)
		if err != nil {
			return c.Status(500).JSON(
				fiber.Map{"err": "Internal server error!"},
			)
		}
		return c.Status(200).JSON(
			fiber.Map{"err": "The folder was created successfully!"},
		)
	} else {
		err = os.WriteFile(fmt.Sprintf("%s%s/%s", config.Folder, itemPath, itemName), nil, 0644)
		if err != nil {
			return c.Status(500).JSON(
				fiber.Map{"err": "Internal server error!"},
			)
		}
		return c.Status(200).JSON(
			fiber.Map{"err": "The file was created successfully!"},
		)
	}
}
