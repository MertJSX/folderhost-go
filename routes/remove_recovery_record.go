package routes

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/MertJSX/folder-host-go/database/logs"
	"github.com/MertJSX/folder-host-go/database/recovery"
	"github.com/MertJSX/folder-host-go/types"
	"github.com/MertJSX/folder-host-go/utils"
	"github.com/MertJSX/folder-host-go/utils/config"
	"github.com/gofiber/fiber/v2"
)

func RemoveRecoveryRecord(c *fiber.Ctx) error {
	if !c.Locals("account").(types.Account).Permissions.UseRecovery {
		return c.Status(403).JSON(
			fiber.Map{"err": "No permission!"},
		)
	}
	var id string = c.Query("id")
	idToInt, err := strconv.Atoi(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"err": "Bad request",
		})
	}

	account := c.Locals("account").(types.Account)
	currentRecord, err := recovery.GetRecoveryRecord(idToInt)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"err": "Error while getting record.",
		})
	}

	var isExistingItem bool = true

	if utils.IsNotExistingPath(currentRecord.BinLocation) {
		isExistingItem = false
	}

	if !strings.HasPrefix(currentRecord.OldLocation, config.Config.Folder+account.Scope) {
		return c.Status(403).JSON(
			fiber.Map{"err": "Out of scope error! No permission!"},
		)
	}

	if isExistingItem {
		if err = os.RemoveAll(currentRecord.BinLocation); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"err": "Error while deleting item.",
			})
		}
	}

	err = recovery.DeleteRecoveryRecord(idToInt, config.Config.Folder+account.Scope)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"err": "Error while deleting database record. But your item was successfully removed.",
		})
	}

	logs.CreateLog(types.AuditLog{
		Username:    c.Locals("account").(types.Account).Username,
		Action:      "Remove Recovery record",
		Description: fmt.Sprintf("%s permanently removed %s recovery record", c.Locals("account").(types.Account).Username, currentRecord.OldLocation),
	})

	return c.Status(200).JSON(fiber.Map{
		"res": "Successfully removed!",
	})
}
