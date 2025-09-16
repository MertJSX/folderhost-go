package routes

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/MertJSX/folder-host-go/types"
	"github.com/MertJSX/folder-host-go/utils"
	"github.com/gofiber/fiber/v2"
)

// Existing problems that I will fix:
// - You can't move the folder or file to recovery_bin if it already exists
func Delete(c *fiber.Ctx) error {

	if !c.Locals("account").(types.Account).Permissions.Delete {
		return c.JSON(
			fiber.Map{"err": "No permission!"},
		)
	}

	config := &utils.Config

	path := fmt.Sprintf("%s%s", config.Folder, c.Query("path"))

	pathStat, err := os.Stat(path)

	if os.IsNotExist(err) {
		return c.JSON(
			fiber.Map{"err": "Wrong path!"},
		)
	}

	if path == fmt.Sprintf("%s/", config.Folder) {
		return c.JSON(
			fiber.Map{"err": "You can't delete the main folder!"},
		)
	}

	if pathStat.IsDir() && !config.RecoveryBin {
		err := os.RemoveAll(path)
		if err == nil {
			return c.Status(200).JSON(fiber.Map{"response": "Item was deleted successfully!"})
		}
	}

	if !config.RecoveryBin {
		err := os.Remove(path)
		if err == nil {
			return c.Status(200).JSON(fiber.Map{"response": "Item was deleted successfully!"})
		}
	}

	itemName := filepath.Base(path)
	_, err = os.Stat(fmt.Sprintf("./recovery_bin/%s%s", itemName, filepath.Ext(path)))

	if os.IsNotExist(err) {
		i := 0
		var err error
		for os.IsNotExist(err) {
			itemName = fmt.Sprintf("%s (%d)%s", filepath.Base(path), i, filepath.Ext(path))
			_, err = os.Stat(fmt.Sprintf("./recovery_bin/%s", itemName))
			i++
		}
	}

	BinStorageLimit := utils.ConvertStringToBytes(config.BinStorageLimit)

	if config.BinStorageLimit != "UNLIMITED" {
		itemToBeDeletedStat, _ := os.Stat(path)
		sizeOfItem := itemToBeDeletedStat.Size()
		if itemToBeDeletedStat.IsDir() {
			sizeOfItem, _, err = utils.GetDirectorySize(path)
			if err != nil {
				return c.Status(500).JSON(fiber.Map{"err": "Unknown server error"})
			}
		}
		sizeOfRecoveryBin, _, err := utils.GetDirectorySize("./recovery_bin")

		if err != nil {
			return c.Status(500).JSON(fiber.Map{"err": "Unknown server error"})
		}

		totalSize := sizeOfRecoveryBin + sizeOfItem

		if totalSize > BinStorageLimit {
			return c.Status(413).JSON(fiber.Map{"err": "This item exceeds the maximum recovery bin size!"})
		}
	}

	err = os.Rename(path, fmt.Sprintf("./recovery_bin/%s", itemName))

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"err": "Error deleting item"})
	}

	return c.Status(200).JSON(fiber.Map{"response": "Item was deleted successfully!"})
}
