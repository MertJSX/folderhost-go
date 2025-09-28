package routes

import (
	"os"
	"strconv"

	"github.com/MertJSX/folder-host-go/database"
	"github.com/MertJSX/folder-host-go/utils"
	"github.com/gofiber/fiber/v2"
)

func RemoveRecoveryRecord(c *fiber.Ctx) error {
	var id string = c.Query("id")
	idToInt, err := strconv.Atoi(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"err": "Bad request",
		})
	}

	currentRecord, err := database.GetRecoveryRecord(idToInt)

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

	err = database.DeleteRecoveryRecord(idToInt)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"err": "Error while deleting database record. But your item was successfully removed.",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"res": "Successfully removed!",
	})
}
