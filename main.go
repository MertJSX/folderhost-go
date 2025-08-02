package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MertJSX/folder-host-go/types"
	"github.com/MertJSX/folder-host-go/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gopkg.in/yaml.v3"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	fileData, err := os.ReadFile("./config.yml")
	if err != nil {
		log.Fatal("Error reading config file:", err)
	}
	var config types.ConfigFile

	err = yaml.Unmarshal(fileData, &config)
	if err != nil {
		log.Fatalf("YAML parse error: %v", err)
	}
	var PORT string = fmt.Sprintf(":%d", config.Port)

	_, dirSize, _ := utils.GetDirectorySize(config.Folder)

	fmt.Printf("%s \n", dirSize)

	utils.GetDirectoryItems(config.Folder, "Test", config)

	app.Static("/", "./client")

	app.Get("*", func(c *fiber.Ctx) error {
		return c.SendFile("./client/index.html")
	})

	app.Listen(PORT)
}
