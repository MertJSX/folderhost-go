package routes

import (
	"fmt"

	"github.com/MertJSX/folder-host-go/database/users"
	"github.com/MertJSX/folder-host-go/types"
	"github.com/gofiber/fiber/v2"
)

func EditUser(c *fiber.Ctx) error {
	if !c.Locals("account").(types.Account).Permissions.EditUsers {
		return c.Status(403).JSON(
			fiber.Map{"err": "No permission!"},
		)
	}

	var requestBody struct {
		User types.Account `json:"user"`
	}

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(400).JSON(
			fiber.Map{"err": "Bad request! " + err.Error()},
		)
	}

	fmt.Println(requestBody.User.ID)

	if requestBody.User.ID == nil {
		return c.Status(400).JSON(fiber.Map{
			"err": "Bad request. User's ID is missing!",
		})
	}

	if *requestBody.User.ID == 1 {
		return c.Status(400).JSON(fiber.Map{
			"err": "You can't update admin account from the web panel. Use config.yml instead.",
		})
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

	err := users.UpdateUser(*requestBody.User.ID, &requestBody.User)

	if err != nil {
		return c.Status(500).JSON(
			fiber.Map{"err": "Unknown server error."},
		)
	}

	return c.Status(200).JSON(
		fiber.Map{"response": "User successfully created!"},
	)
}
