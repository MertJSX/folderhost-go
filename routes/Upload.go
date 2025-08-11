package routes

import (
	"fmt"
	"os"

	"github.com/MertJSX/folder-host-go/utils"
	"github.com/gofiber/fiber/v2"
)

func Upload(c *fiber.Ctx) error {
	config := utils.GetConfig()
	targetPath := c.Query("path", "")
	if targetPath == "" {
		return c.Status(400).JSON(fiber.Map{
			"err": "Missing path query",
		})
	}

	file, err := c.FormFile("file")
	fmt.Printf("Filesize of uploaded file is: %s\n", utils.ConvertBytesToString(file.Size))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"err": "File is missing",
		})
	}

	savePath := fmt.Sprintf("%s%s", config.Folder, targetPath)

	fileinfo, err := os.Stat(savePath)

	// Validation to avoid errors
	if os.IsNotExist(err) {
		return c.JSON(
			fiber.Map{"err": "Wrong path!"},
		)
	} else if !fileinfo.IsDir() {
		return c.JSON(
			fiber.Map{"err": "Path is targeting a file!"},
		)
	}

	fullSavePath := fmt.Sprintf("%s%s", savePath, file.Filename)

	if err := c.SaveFile(file, fullSavePath); err != nil {
		fmt.Printf("Error while uploading file:\n %s\n", err)
		return c.Status(500).JSON(fiber.Map{
			"err": "Error uploading file",
		})
	}

	return c.JSON(fiber.Map{
		"response": "Successfully uploaded!",
	})
}
