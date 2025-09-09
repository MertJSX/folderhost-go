package middleware

import (
	"log"

	"github.com/MertJSX/folder-host-go/utils"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func WsConnect(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	var username string = ""
	var err error

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token required",
		})
	}

	username, err = utils.VerifyToken(token, utils.Config.SecretJwtKey)

	if err != nil {
		return c.Status(401).JSON(fiber.Map{"err": "Invalid token"})
	}

	c.Locals("username", username)

	config := &utils.Config
	for _, v := range config.Accounts {
		if v.Name != username {
			continue
		}

		c.Locals("account", v)
		return c.Next()
	}

	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		c.Locals("token", token)
		return c.Next()
	}

	return fiber.ErrUpgradeRequired
}

func HandleWebsocket(c *websocket.Conn) {
	log.Println(c.Locals("allowed")) // true
	log.Println(c.Params("path"))    // /example

	// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
	var (
		mt  int
		msg []byte
		err error
	)
	for {
		if mt, msg, err = c.ReadMessage(); err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", msg)

		if err = c.WriteMessage(mt, msg); err != nil {
			log.Println("write:", err)
			break
		}
	}
}
