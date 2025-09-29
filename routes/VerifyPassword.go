package routes

import (
	"github.com/MertJSX/folder-host-go/types"
	"github.com/MertJSX/folder-host-go/utils"
	"github.com/gofiber/fiber/v2"
)

func VerifyPassword(c *fiber.Ctx) error {
	token, err := utils.CreateToken(c.Locals("account").(types.Account).Username, utils.Config.SecretJwtKey)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"err": "unknown error while getting token"})
	}

	return c.JSON(
		fiber.Map{
			"res":         "Verified!",
			"token":       token,
			"permissions": c.Locals("account").(types.Account).Permissions,
		},
	)
}
