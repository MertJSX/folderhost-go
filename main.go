package main

import (
	"fmt"

	"github.com/MertJSX/folder-host-go/database"
	"github.com/MertJSX/folder-host-go/middleware"
	_ "github.com/MertJSX/folder-host-go/resources"
	"github.com/MertJSX/folder-host-go/routes"
	"github.com/MertJSX/folder-host-go/utils"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 1000 * 1024 * 1024, // 1 GB
	})
	app.Use(cors.New())

	utils.Setup()
	utils.GetConfig()

	database.InitializeDatabase()

	var PORT string = fmt.Sprintf(":%d", utils.Config.Port)

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return middleware.WsConnect(c)
		}
		return c.Next()
	})

	app.Get("/ws/:path", websocket.New(func(c *websocket.Conn) {
		middleware.HandleWebsocket(c)
	}))

	app.Use("/api", func(c *fiber.Ctx) error {
		return middleware.CheckAuth(c)
	})

	app.Get("/api/read-file", func(c *fiber.Ctx) error {
		return routes.ReadFile(c)
	})

	app.Post("/api/verify-password", func(c *fiber.Ctx) error {
		return routes.VerifyPassword(c)
	})

	app.Get("/api/read-dir", func(c *fiber.Ctx) error {
		return routes.ReadDirectory(c)
	})

	app.Get("/api/download", func(c *fiber.Ctx) error {
		return routes.Download(c)
	})

	app.Post("/api/upload", func(c *fiber.Ctx) error {
		return routes.ChunkedUpload(c)
	})

	app.Get("/api/delete", func(c *fiber.Ctx) error {
		return routes.Delete(c)
	})

	app.Get("/api/create-item", func(c *fiber.Ctx) error {
		return routes.CreateItem(c)
	})

	app.Get("/api/create-copy", func(c *fiber.Ctx) error {
		return routes.CreateCopy(c)
	})

	app.Get("/api/rename", func(c *fiber.Ctx) error {
		return routes.Rename(c)
	})

	app.Get("/api/recovery", func(c *fiber.Ctx) error {
		return routes.Recovery(c)
	})

	app.Get("/api/recover-item", func(c *fiber.Ctx) error {
		return routes.RecoverItem(c)
	})

	app.Static("/", "client/dist")

	app.Get("*", func(c *fiber.Ctx) error {
		return c.SendFile("client/dist/index.html")
	})

	app.Listen(PORT)
}
