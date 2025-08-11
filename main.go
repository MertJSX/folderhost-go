package main

import (
	"fmt"

	"github.com/MertJSX/folder-host-go/middleware"
	_ "github.com/MertJSX/folder-host-go/resources"
	"github.com/MertJSX/folder-host-go/routes"
	"github.com/MertJSX/folder-host-go/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 1000 * 1024 * 1024, // 1 GB
	})
	app.Use(cors.New())

	config := utils.GetConfig()

	var PORT string = fmt.Sprintf(":%d", config.Port)

	app.Use("/api", func(c *fiber.Ctx) error {
		return middleware.CheckAuth(c)
	})

	app.Post("/api/verify-password", func(c *fiber.Ctx) error {
		return routes.VerifyPassword(c)
	})

	app.Post("/api/read-dir", func(c *fiber.Ctx) error {
		return routes.ReadDirectory(c)
	})

	app.Post("/api/download", func(c *fiber.Ctx) error {
		return routes.Download(c)
	})

	app.Post("/api/upload", func(c *fiber.Ctx) error {
		return routes.Upload(c)
	})

	app.Static("/", "client")

	app.Get("*", func(c *fiber.Ctx) error {
		return c.SendFile("client/index.html")
	})

	app.Listen(PORT)
}
