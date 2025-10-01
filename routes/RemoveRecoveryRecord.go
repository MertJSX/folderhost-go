package routes

import (
	"os"
	"strconv"

	"github.com/MertJSX/folder-host-go/database/recovery"
	"github.com/MertJSX/folder-host-go/types"
	"github.com/MertJSX/folder-host-go/utils"
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

	if isExistingItem {
		if err = os.RemoveAll(currentRecord.BinLocation); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"err": "Error while deleting item.",
			})
		}
	}

	err = recovery.DeleteRecoveryRecord(idToInt)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"err": "Error while deleting database record. But your item was successfully removed.",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"res": "Successfully removed!",
	})
}
