package routes

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/MertJSX/folder-host-go/types"
	"github.com/MertJSX/folder-host-go/utils"
	"github.com/gofiber/fiber/v2"
)

func Rename(c *fiber.Ctx) error {
	config := &utils.Config
	oldFilepath := c.Query("oldFilepath")
	var filename string
	newFilepath := c.Query("newFilepath")
	requestType := c.Query("type")

	// Check permissions
	if requestType == "rename" && !c.Locals("account").(types.Account).Permissions.Rename {
		return c.Status(403).JSON(
			fiber.Map{"err": "No permission!"},
		)
	} else if requestType == "move" && !c.Locals("account").(types.Account).Permissions.Move {
		return c.Status(403).JSON(
			fiber.Map{"err": "No permission!"},
		)
	}

	if requestType != "move" && requestType != "rename" {
		return c.Status(400).JSON(fiber.Map{"err": "Bad request!"})
	}

	if requestType == "move" {
		newFilepathStat, err := os.Stat(fmt.Sprintf("%s%s", config.Folder, newFilepath))
		if os.IsNotExist(err) {
			return c.Status(400).JSON(fiber.Map{"err": "Newpath is not existing!"})
		}
		if !newFilepathStat.IsDir() {
			return c.Status(400).JSON(fiber.Map{"err": "Newpath is not directory!"})
		}
	}

	if oldFilepath == newFilepath {
		return c.Status(400).JSON(fiber.Map{"err": "Same location!"})
	}

	if utils.IsNotExistingPath(fmt.Sprintf("%s%s", config.Folder, oldFilepath)) {
		return c.Status(400).JSON(fiber.Map{"err": "Filepath doesn't exist!"})
	}

	if requestType == "rename" && utils.IsNotExistingPath(fmt.Sprintf("%s%s", config.Folder, utils.GetParentPath(newFilepath))) {
		return c.Status(400).JSON(fiber.Map{"err": "New parent directory doesn't exist!"})
	}

	filename = filepath.Base(oldFilepath)

	if requestType == "move" {
		// Check possible existing item in the new directory with the same name
		oldPathPlaceholder := fmt.Sprintf("%s%s", config.Folder, oldFilepath)
		newPathPlaceholder := fmt.Sprintf("%s%s/%s", config.Folder, newFilepath, filename)
		fmt.Printf("Old path: %s\n", oldPathPlaceholder)
		fmt.Printf("New path: %s\n", newPathPlaceholder)
		if !utils.IsNotExistingPath(newPathPlaceholder) {
			return c.Status(500).JSON(fiber.Map{"err": "The destination already has an item named like that!"})
		}

		err := os.Rename(oldPathPlaceholder, newPathPlaceholder)

		if err != nil {
			return c.Status(520).JSON(fiber.Map{"err": "Unknown error while moving item"})
		}
	} else {
		oldPathPlaceholder := fmt.Sprintf("%s/%s", config.Folder, filename)
		newPathPlaceholder := fmt.Sprintf("%s%s", config.Folder, newFilepath)
		if !utils.IsNotExistingPath(newPathPlaceholder) {
			return c.Status(500).JSON(fiber.Map{"err": "The destination already has an item named!"})
		}

		err := os.Rename(oldPathPlaceholder, newPathPlaceholder)

		if err != nil {
			fmt.Printf("Error while renaming item: %s", err)
			return c.Status(520).JSON(fiber.Map{"err": "Unknown error while renaming item"})
		}
	}

	return c.Status(200).JSON(fiber.Map{"response": "Saved!"})
}
