package routes

import (
	"github.com/MertJSX/folder-host-go/database/users"
	"github.com/MertJSX/folder-host-go/types"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	if !c.Locals("account").(types.Account).Permissions.EditUsers {
		return c.Status(403).JSON(
			fiber.Map{"err": "No permission!"},
		)
	}

	var requestBody struct {
		User types.Account `json:"users"`
	}

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(400).JSON(
			fiber.Map{"err": "Bad request! " + err.Error()},
		)
	}

	if requestBody.User.Username == "" {
		return c.Status(400).JSON(
			fiber.Map{"err": "Username is missing."},
		)
	}

	if requestBody.User.Password == "" {
		return c.Status(400).JSON(
			fiber.Map{"err": "Password is missing."},
		)
	}

	err := users.CreateUser(&requestBody.User)

	if err.Error() == "username already exists" {
		return c.Status(400).JSON(
			fiber.Map{"err": "Username already exists."},
		)
	}

	if err != nil {
		return c.Status(500).JSON(
			fiber.Map{"err": "Unknown server error."},
		)
	}

	return c.Status(200).JSON(
		fiber.Map{"response": "User successfully created!"},
	)
}
