package routes

import (
	"github.com/MertJSX/folder-host-go/database/recovery"
	"github.com/MertJSX/folder-host-go/types"
	"github.com/MertJSX/folder-host-go/utils"
	"github.com/gofiber/fiber/v2"
)

func ResetRecoveryRecords(c *fiber.Ctx) error {
	if !c.Locals("account").(types.Account).Permissions.UseRecovery {
		return c.Status(403).JSON(
			fiber.Map{"err": "No permission!"},
		)
	}

	if err := utils.ClearDirectory("./recovery_bin"); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"err": "Error while clearing recovery bin.",
		})
	}

	err := recovery.ResetRecoveryRecords()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"err": "Error while deleting database records. But your items were successfully removed.",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"res": "Successfully cleared!",
	})
}
