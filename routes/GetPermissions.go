package routes

import (
	"github.com/MertJSX/folder-host-go/types"
	"github.com/gofiber/fiber/v2"
)

func GetPermissions(c *fiber.Ctx) error {

	return c.JSON(
		fiber.Map{
			"permissions": c.Locals("account").(types.Account).Permissions,
		},
	)
}
