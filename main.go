package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/MertJSX/folder-host-go/database/initialize"
	"github.com/MertJSX/folder-host-go/middleware"
	fhWS "github.com/MertJSX/folder-host-go/middleware/websocket"
	_ "github.com/MertJSX/folder-host-go/resources"
	"github.com/MertJSX/folder-host-go/routes"
	"github.com/MertJSX/folder-host-go/types"
	"github.com/MertJSX/folder-host-go/utils"
	"github.com/MertJSX/folder-host-go/utils/cache"
	"github.com/MertJSX/folder-host-go/utils/tasks"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed client/dist/*
var FrontendFS embed.FS

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 10 * 1024 * 1024, // 10 MB
	})

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
		Next: func(c *fiber.Ctx) bool {
			skipRoutes := []string{
				"/api/explorer/download",
				"/api/upload",
			}
			for _, route := range skipRoutes {
				if c.Path() == route {
					return true
				}
			}
			return false
		},
	}))

	app.Use(cors.New())

	utils.Setup()
	utils.GetConfig()
	initialize.InitializeDatabase()

	go cache.ListenDirectorySetCacheEvents()
	go tasks.AutoClearOldLogs()

	var portInt int = utils.Config.Port
	if portInt == 0 {
		portInt = 5000
	}
	var PORT string = fmt.Sprintf(":%d", portInt)

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return fhWS.WsConnect(c)
		}
		return c.Next()
	})

	app.Get("/ws/:path", websocket.New(func(c *websocket.Conn) {
		fhWS.HandleWebsocket(c)
	}))

	app.Use("/api", func(c *fiber.Ctx) error {
		return middleware.CheckAuth(c)
	})

	app.Get("/api/user-info", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"username": c.Locals("account").(types.Account).Username,
		})
	})

	app.Get("/api/read-file", func(c *fiber.Ctx) error {
		return routes.ReadFile(c)
	})

	app.Post("/api/verify-password", func(c *fiber.Ctx) error {
		return routes.VerifyPassword(c)
	})

	app.Get("/api/permissions", func(c *fiber.Ctx) error {
		return routes.GetPermissions(c)
	})

	app.Get("/api/explorer/read-dir", func(c *fiber.Ctx) error {
		return routes.ReadDirectory(c)
	})

	app.Get("/api/explorer/download", func(c *fiber.Ctx) error {
		return routes.Download(c)
	})

	app.Post("/api/upload", func(c *fiber.Ctx) error {
		return routes.ChunkedUpload(c)
	})

	app.Delete("/api/explorer/delete", func(c *fiber.Ctx) error {
		return routes.Delete(c)
	})

	app.Post("/api/explorer/create-item", func(c *fiber.Ctx) error {
		return routes.CreateItem(c)
	})

	app.Post("/api/explorer/create-copy", func(c *fiber.Ctx) error {
		return routes.CreateCopy(c)
	})

	app.Put("/api/explorer/rename", func(c *fiber.Ctx) error {
		return routes.Rename(c)
	})

	app.Get("/api/recovery", func(c *fiber.Ctx) error {
		return routes.Recovery(c)
	})

	app.Put("/api/recovery/recover", func(c *fiber.Ctx) error {
		return routes.RecoverItem(c)
	})

	app.Delete("/api/recovery/remove", func(c *fiber.Ctx) error {
		return routes.RemoveRecoveryRecord(c)
	})

	app.Delete("/api/recovery/clear", func(c *fiber.Ctx) error {
		return routes.ResetRecoveryRecords(c)
	})

	app.Get("/api/users", func(c *fiber.Ctx) error {
		return routes.GetAllUsers(c)
	})

	app.Get("/api/users/:username", func(c *fiber.Ctx) error {
		return routes.GetUser(c)
	})

	app.Put("/api/users/edit", func(c *fiber.Ctx) error {
		return routes.EditUser(c)
	})

	app.Post("/api/users/new", func(c *fiber.Ctx) error {
		return routes.CreateUser(c)
	})

	app.Delete("/api/users/remove/:id", func(c *fiber.Ctx) error {
		return routes.RemoveUser(c)
	})

	app.Get("/api/logs", func(c *fiber.Ctx) error {
		return routes.Logs(c)
	})

	if !utils.IsDevelopment() {
		distFS, err := fs.Sub(FrontendFS, "client/dist")
		if err != nil {
			log.Fatal("Error creating sub FS:", err)
		}

		app.Use("/", filesystem.New(filesystem.Config{
			Root:         http.FS(distFS),
			Index:        "index.html",
			NotFoundFile: "index.html",
			MaxAge:       86400, // 1 day cache
		}))

		app.Get("*", func(c *fiber.Ctx) error {
			indexFile, err := distFS.Open("index.html")
			if err != nil {
				return c.Status(404).SendString("Not Found")
			}
			defer indexFile.Close()

			stat, err := indexFile.Stat()
			if err != nil {
				return c.Status(500).SendString("Internal Server Error")
			}

			c.Type("html")
			return c.SendStream(indexFile, int(stat.Size()))
		})
	}

	if err := app.Listen(PORT); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
