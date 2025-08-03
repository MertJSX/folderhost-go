package middleware

import (
	"github.com/MertJSX/folder-host-go/utils"
	"github.com/gofiber/fiber/v2"
)

func CheckAuth(c *fiber.Ctx) error {
	var body map[string]interface{}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"err": "Bad request"})
	}

	if body["username"].(string) == "" || body["password"].(string) == "" {
		if body["token"].(string) == "" || c.Get("token") == "" {
			return c.JSON(
				fiber.Map{"err": "Bad request!"},
			)
		}
	}

	if body["token"].(string) != "" {
		username, _ := utils.VerifyToken(body["token"].(string), utils.GetConfig().SecretJwtKey)

		c.Locals("username", username)
	}

	return c.Next()
}
