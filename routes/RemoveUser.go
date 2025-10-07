package routes

import (
	"strconv"

	"github.com/MertJSX/folder-host-go/database/users"
	"github.com/MertJSX/folder-host-go/types"
	"github.com/MertJSX/folder-host-go/utils/cache"
	"github.com/gofiber/fiber/v2"
)

func RemoveUser(c *fiber.Ctx) error {
	if !c.Locals("account").(types.Account).Permissions.EditUsers {
		return c.Status(403).JSON(
			fiber.Map{"err": "No permission!"},
		)
	}
	var id string = c.Params("id")
	idToInt, err := strconv.Atoi(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"err": "Bad request",
		})
	}

	if idToInt == 1 {
		return c.Status(403).JSON(fiber.Map{
			"err": "You can't remove the admin account!",
		})
	}

	username, err := users.GetUsername(idToInt)

	if err != nil {
		cache.SessionCache.Clear()
	}

	cache.SessionCache.Delete(username)

	err = users.RemoveUser(idToInt)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"err": "Internal server error",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"res": "Successfully removed!",
	})
}
