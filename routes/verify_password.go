package routes

import (
	"fmt"

	"github.com/MertJSX/folder-host-go/database/logs"
	"github.com/MertJSX/folder-host-go/types"
	"github.com/MertJSX/folder-host-go/utils"
	"github.com/MertJSX/folder-host-go/utils/config"
	"github.com/gofiber/fiber/v2"
)

func VerifyPassword(c *fiber.Ctx) error {
	token, err := utils.CreateToken(c.Locals("account").(types.Account).Username, config.Config.SecretJwtKey)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"err": "unknown error while getting token"})
	}

	logs.CreateLog(types.AuditLog{
		Username:    c.Locals("account").(types.Account).Username,
		Action:      "Login",
		Description: fmt.Sprintf("%s logged in to his account.", c.Locals("account").(types.Account).Username),
	})

	return c.JSON(
		fiber.Map{
			"res":         "Verified!",
			"token":       token,
			"permissions": c.Locals("account").(types.Account).Permissions,
		},
	)
}
