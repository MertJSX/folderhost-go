package routes

import (
	"github.com/MertJSX/folder-host-go/database/users"
	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {

	users, err := users.GetAll()

	if err != nil {
		return c.Status(500).JSON(
			fiber.Map{"err": "Unknown error!"},
		)
	}

	return c.Status(200).JSON(
		fiber.Map{"users": users},
	)
}
