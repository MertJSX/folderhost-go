package middleware

import (
	"github.com/MertJSX/folder-host-go/utils"
	"github.com/gofiber/fiber/v2"
)

func CheckAuth(c *fiber.Ctx) error {
	var body map[string]interface{}
	var controlPassword bool = false
	var username string = ""
	var password string = ""
	var err error

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"err": "Bad request"})
	}

	token, hasToken := body["token"].(string)
	if !hasToken {
		token = c.Get("token")
	}

	reqUsername, hasUsername := body["username"].(string)
	reqPassword, hasPassword := body["password"].(string)

	if token == "" && (!hasUsername || !hasPassword || reqUsername == "" || reqPassword == "") {
		return c.Status(400).JSON(fiber.Map{"err": "Bad request! Authentication required"})
	}

	if token != "" {
		username, err = utils.VerifyToken(token, utils.GetConfig().SecretJwtKey)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{"err": "Invalid token"})
		}
		c.Locals("username", username)
	} else {
		username = reqUsername
		password = reqPassword
		controlPassword = true
	}

	accounts := utils.GetConfig().Accounts
	for _, v := range accounts {
		if v.Name != username {
			continue
		}
		if controlPassword && password != v.Password {
			return c.Status(401).JSON(fiber.Map{"err": "Wrong password!"})
		}

		c.Locals("account", v)
		return c.Next()
	}

	return c.Status(404).JSON(fiber.Map{"err": "Account not found"})
}
