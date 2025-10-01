package routes

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/MertJSX/folder-host-go/database/recovery"
	"github.com/MertJSX/folder-host-go/types"
	"github.com/MertJSX/folder-host-go/utils"
	"github.com/gofiber/fiber/v2"
)

func Delete(c *fiber.Ctx) error {

	if !c.Locals("account").(types.Account).Permissions.Delete {
		return c.Status(403).JSON(
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
		} else {
			return c.Status(500).JSON(fiber.Map{"err": "Unknown error!"})
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

	itemToBeDeletedStat, _ := os.Stat(path)
	isDirectory := itemToBeDeletedStat.IsDir()
	sizeOfItem := itemToBeDeletedStat.Size()
	if itemToBeDeletedStat.IsDir() {
		sizeOfItem, _, err = utils.GetDirectorySize(path)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"err": "Unknown server error"})
		}
	}

	if config.BinStorageLimit != "UNLIMITED" {
		sizeOfRecoveryBin, _, err := utils.GetDirectorySize("./recovery_bin")

		if err != nil {
			return c.Status(500).JSON(fiber.Map{"err": "Unknown server error"})
		}

		totalSize := sizeOfRecoveryBin + sizeOfItem

		if totalSize > BinStorageLimit {
			return c.Status(413).JSON(fiber.Map{"err": "This item exceeds the maximum recovery bin size!"})
		}
	}

	var (
		originalName string = utils.GetPureFileName(itemName)
		extName      string = filepath.Ext(itemName)
		copyIndex    int    = 0
		fullFileName string = originalName + extName
	)

	if isDirectory {
		fullFileName = itemName
		for utils.IsExistingPath(fmt.Sprintf("./recovery_bin/%s", fullFileName)) {
			copyIndex++
			fullFileName = fmt.Sprintf("%s (%d)", fullFileName, copyIndex)
		}
	} else {
		for utils.IsExistingPath(fmt.Sprintf("./recovery_bin/%s", fullFileName)) {
			copyIndex++
			fullFileName = fmt.Sprintf("%s (%d)%s", originalName, copyIndex, extName)
		}
	}

	err = os.Rename(path, fmt.Sprintf("./recovery_bin/%s", fullFileName))

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"err": "Error deleting item"})
	}

	var recoveryRecord types.RecoveryRecord = types.RecoveryRecord{
		Username:    c.Locals("account").(types.Account).Username,
		OldLocation: path,
		BinLocation: fmt.Sprintf("./recovery_bin/%s", fullFileName),
		IsDirectory: isDirectory,
		SizeDisplay: utils.ConvertBytesToString(sizeOfItem),
		SizeBytes:   sizeOfItem,
	}

	if err = recovery.CreateRecoveryRecord(recoveryRecord); err != nil {
		fmt.Printf("Error: %s", err)
		return c.Status(500).JSON(fiber.Map{"err": "An error occurred during the creation of the recovery record. But the item was moved to the recovery bin."})
	}

	return c.Status(200).JSON(fiber.Map{"response": "Item was deleted successfully!"})
}
