package routes

import (
	"fmt"

	"github.com/MertJSX/folder-host-go/database/users"
	"github.com/MertJSX/folder-host-go/types"
	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	if !c.Locals("account").(types.Account).Permissions.ReadUsers {
		return c.Status(403).JSON(
			fiber.Map{"err": "No permission!"},
		)
	}

	var username string = c.Params("username")

	fmt.Println(username)

	if username == "" {
		return c.Status(400).JSON(
			fiber.Map{"err": "Username is missing!"},
		)
	}

	isExist, err := users.CheckIfUsernameExists(username)

	if err != nil {
		return c.Status(500).JSON(
			fiber.Map{"err": "Internal server error!"},
		)
	}

	if !isExist {
		return c.Status(400).JSON(
			fiber.Map{"err": "Username does not exist!"},
		)
	}

	user, err := users.GetUserByUsername(username)

	if err != nil {
		return c.Status(500).JSON(
			fiber.Map{"err": "Unknown error!"},
		)
	}

	return c.Status(200).JSON(
		fiber.Map{"user": user},
	)
}
