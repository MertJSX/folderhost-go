package websocket

import (
	"strings"

	"github.com/MertJSX/folder-host-go/database/users"
	"github.com/MertJSX/folder-host-go/utils"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func WsConnect(c *fiber.Ctx) error {
	if !websocket.IsWebSocketUpgrade(c) {
		return c.Next()
	}

	token := c.Query("token")
	if token == "" {
		authHeader := c.Get("Authorization")
		if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
			token = strings.TrimPrefix(authHeader, "Bearer ")
		}
	}

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token required for WebSocket connection",
		})
	}

	username, err := utils.VerifyToken(token, utils.Config.SecretJwtKey)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "invalid token"})
	}

	foundAccount, err := users.GetUserByUsername(username)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"err": "account not found"})
	}

	c.Locals("username", foundAccount.Username)
	c.Locals("account", foundAccount)
	c.Locals("token", token)
	c.Locals("allowed", true)

	return c.Next()
}
