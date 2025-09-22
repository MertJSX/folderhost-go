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
	var token string

	c.BodyParser(&body)

	token, hasToken := body["token"].(string)
	path := c.Query("path")
	folder := c.Query("folder")
	itemName := c.Query("itemName")
	filepath := c.Query("filepath")
	oldFilepath := c.Query("oldFilepath")
	newFilepath := c.Query("newFilepath")

	if path != "" && !utils.IsSafePath(path) {
		return c.Status(403).JSON(fiber.Map{"err": "Forbidden."})
	}

	if folder != "" && !utils.IsSafePath(folder) {
		return c.Status(403).JSON(fiber.Map{"err": "Forbidden."})
	}

	if itemName != "" && !utils.IsSafePath(itemName) {
		return c.Status(403).JSON(fiber.Map{"err": "Forbidden."})
	}

	if filepath != "" && !utils.IsSafePath(filepath) {
		return c.Status(403).JSON(fiber.Map{"err": "Forbidden."})
	}

	if oldFilepath != "" && !utils.IsSafePath(oldFilepath) {
		return c.Status(403).JSON(fiber.Map{"err": "Forbidden."})
	}

	if newFilepath != "" && !utils.IsSafePath(newFilepath) {
		return c.Status(403).JSON(fiber.Map{"err": "Forbidden."})
	}

	if !hasToken {
		token = c.Get("token")
	}

	reqUsername, hasUsername := body["username"].(string)
	reqPassword, hasPassword := body["password"].(string)

	if token == "" && (!hasUsername || !hasPassword || reqUsername == "" || reqPassword == "") {
		return c.Status(400).JSON(fiber.Map{"err": "Bad request! Authentication required"})
	}

	if token != "" {
		username, err = utils.VerifyToken(token, utils.Config.SecretJwtKey)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{"err": "Invalid token"})
		}
		c.Locals("username", username)
	} else {
		username = reqUsername
		password = reqPassword
		controlPassword = true
	}

	config := &utils.Config
	for _, v := range config.Accounts {
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
