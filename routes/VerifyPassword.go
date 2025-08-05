package routes

import (
	"github.com/MertJSX/folder-host-go/types"
	"github.com/MertJSX/folder-host-go/utils"
	"github.com/gofiber/fiber/v2"
)

func VerifyPassword(c *fiber.Ctx) error {
	config := utils.GetConfig()

	for i := range len(config.Accounts) {
		if c.Locals("account").(types.Account).Name == config.Accounts[i].Name {
			if c.Locals("account").(types.Account).Password == config.Accounts[i].Password {
				token, _ := utils.CreateToken(config.Accounts[i].Name, config.SecretJwtKey)
				return c.JSON(
					fiber.Map{
						"res":         "Verified!",
						"token":       token,
						"permissions": config.Accounts[i].Permissions,
					},
				)
			} else {
				return c.JSON(
					fiber.Map{"err": "Incorrect password!"},
				)
			}
		}
	}

	return c.JSON(
		fiber.Map{"err": "Username was not found!"},
	)
}
