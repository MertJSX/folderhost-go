package routes

import (
	"fmt"

	"github.com/MertJSX/folder-host-go/database/logs"
	"github.com/MertJSX/folder-host-go/database/recovery"
	"github.com/MertJSX/folder-host-go/types"
	"github.com/MertJSX/folder-host-go/utils"
	"github.com/MertJSX/folder-host-go/utils/config"
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

	account := c.Locals("account").(types.Account)
	err := recovery.ResetRecoveryRecords(config.Config.Folder + account.Scope)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"err": "Error while deleting database records. But your items were successfully removed.",
		})
	}

	logs.CreateLog(types.AuditLog{
		Username:    c.Locals("account").(types.Account).Username,
		Action:      "Clear Recovery",
		Description: fmt.Sprintf("%s cleared all the recovery records", c.Locals("account").(types.Account).Username),
	})

	return c.Status(200).JSON(fiber.Map{
		"res": "Successfully cleared!",
	})
}
