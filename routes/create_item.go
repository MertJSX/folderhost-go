package routes

import (
	"fmt"
	"os"
	"strconv"

	"github.com/MertJSX/folder-host-go/database/logs"
	"github.com/MertJSX/folder-host-go/types"
	"github.com/MertJSX/folder-host-go/utils"
	"github.com/MertJSX/folder-host-go/utils/config"
	"github.com/gofiber/fiber/v2"
)

func CreateItem(c *fiber.Ctx) error {
	var (
		itemPath string            = c.Query("path")
		itemName string            = c.Query("itemName")
		account  types.Account     = c.Locals("account").(types.Account)
		config   *types.ConfigFile = &config.Config
		isFolder bool
		scope    string = c.Locals("account").(types.Account).Scope
	)

	if !account.Permissions.Create {
		return c.Status(403).JSON(
			fiber.Map{"err": "No permission!"},
		)
	}

	if utils.IsExistingPath(fmt.Sprintf("%s%s/%s", config.GetScopedFolder(scope), itemPath, itemName)) {
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
		err = os.Mkdir(fmt.Sprintf("%s%s/%s", config.GetScopedFolder(scope), itemPath, itemName), 0777)
		if err != nil {
			return c.Status(500).JSON(
				fiber.Map{"err": "Internal server error!"},
			)
		}

		logs.CreateLog(types.AuditLog{
			Username:    c.Locals("account").(types.Account).Username,
			Action:      "Create folder",
			Description: fmt.Sprintf("%s created a %s%s folder.", c.Locals("account").(types.Account).Username, itemPath, itemName),
		})

		return c.Status(200).JSON(
			fiber.Map{"err": "The folder was created successfully!"},
		)
	} else {
		err = os.WriteFile(fmt.Sprintf("%s%s/%s", config.GetScopedFolder(scope), itemPath, itemName), nil, 0777)
		if err != nil {
			return c.Status(500).JSON(
				fiber.Map{"err": "Internal server error!"},
			)
		}

		logs.CreateLog(types.AuditLog{
			Username:    c.Locals("account").(types.Account).Username,
			Action:      "Create file",
			Description: fmt.Sprintf("%s created a %s%s file", c.Locals("account").(types.Account).Username, itemPath, itemName),
		})

		return c.Status(200).JSON(
			fiber.Map{"err": "The file was created successfully!"},
		)
	}
}
