package routes

import (
	"fmt"
	"os"
	"strconv"

	"github.com/MertJSX/folder-host-go/database/logs"
	"github.com/MertJSX/folder-host-go/database/recovery"
	"github.com/MertJSX/folder-host-go/types"
	"github.com/MertJSX/folder-host-go/utils"
	"github.com/gofiber/fiber/v2"
)

// Missing: Check for space requirements
func RecoverItem(c *fiber.Ctx) error {
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

	if utils.IsExistingPath(currentRecord.OldLocation) {
		return c.Status(400).JSON(fiber.Map{
			"err": "There is existing item with the same name.",
		})
	}

	if utils.Config.StorageLimit != "UNLIMITED" {
		remainingFreeSpace, err := utils.GetRemainingFolderSpace()

		if err != nil {
			return c.Status(520).JSON(fiber.Map{"err": "Internal server error!"})
		}

		if currentRecord.SizeBytes > remainingFreeSpace {
			return c.Status(413).JSON(fiber.Map{"err": "This item exceeds the storage limit!"})
		}
	}

	if err = os.Rename(currentRecord.BinLocation, currentRecord.OldLocation); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"err": "Error while moving item.",
		})
	}

	if utils.IsNotExistingPath(currentRecord.OldLocation) {
		return c.Status(500).JSON(fiber.Map{
			"err": "Unknown error! Moved item is not in the right place.",
		})
	}

	err = recovery.DeleteRecoveryRecord(idToInt)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"err": "Error while deleting useless database record. But your item was successfully recovered.",
		})
	}

	logs.CreateLog(types.AuditLog{
		Username:    c.Locals("account").(types.Account).Username,
		Action:      "Recover record",
		Description: fmt.Sprintf("%s recovered %s", c.Locals("account").(types.Account).Username, currentRecord.OldLocation),
	})

	return c.Status(200).JSON(fiber.Map{
		"res": "Successfully recovered!",
	})
}
