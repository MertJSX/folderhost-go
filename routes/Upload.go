package routes

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/MertJSX/folder-host-go/types"
	"github.com/MertJSX/folder-host-go/utils"
	"github.com/gofiber/fiber/v2"
)

func Upload(c *fiber.Ctx) error {
	config := &utils.Config
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

func ChunkedUpload(c *fiber.Ctx) error {
	if !c.Locals("account").(types.Account).Permissions.UploadFiles {
		return c.Status(403).JSON(
			fiber.Map{"err": "No permission!"},
		)
	}
	config := &utils.Config
	targetPath := c.Query("path")
	if targetPath == "" {
		return c.Status(400).JSON(fiber.Map{
			"err": "Missing path query",
		})
	}

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(500).SendString("Couldn't read form: " + err.Error())
	}
	defer form.RemoveAll()

	fileID := c.FormValue("fileID")
	chunkIndex := c.FormValue("chunkIndex")
	totalChunks := c.FormValue("totalChunks")
	fileName := c.FormValue("fileName")
	total, _ := strconv.ParseInt(totalChunks, 10, 64)
	// fileSize := total * form.File["file"][0].Size

	// mainFolderSize := utils.GetDirectorySizeAsync()

	// Save chunk as temp file
	chunkPath := filepath.Join("./tmp", fileID+"_"+chunkIndex)
	chunkFile, err := form.File["file"][0].Open()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"err": "Couldn't open chunk",
		})
	}
	defer chunkFile.Close()

	outFile, err := os.Create(chunkPath)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"err": "Couldn't create chunk file",
		})
	}
	defer outFile.Close()

	chunkContent, _ := utils.FileToString(chunkFile)

	ch := make(chan error, 1)
	var wg sync.WaitGroup

	wg.Add(1)
	go utils.CreateFileAsync(chunkPath, chunkContent, &wg, ch)

	go func() {
		wg.Wait()
		close(ch)
	}()

	err = <-ch

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"err": "Couldn't save chunk",
		})
	}

	// Merge all chunks
	currentChunk, _ := strconv.Atoi(chunkIndex)
	if currentChunk == int(total)-1 { // If it's the last chunk
		finalPath := filepath.Join(config.Folder, targetPath, fileName)
		if err := mergeChunks(fileID, finalPath, int(total)); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"err": "Error uploading file",
			})
		}
		return c.JSON(fiber.Map{
			"response": "Successfully uploaded!",
			"uploaded": true,
		})
	}

	return c.JSON(fiber.Map{
		"response": fmt.Sprintf("Uploaded chunk %s", chunkIndex),
	})
}

func mergeChunks(fileID, outputPath string, totalChunks int) error {
	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	for i := 0; i < totalChunks; i++ {
		chunkPath := filepath.Join("./tmp", fmt.Sprintf("%s_%d", fileID, i))
		chunkData, err := os.Open(chunkPath)
		if err != nil {
			return fmt.Errorf("couldn't find chunk %d: %v", i, err)
		}

		if _, err := io.Copy(outFile, chunkData); err != nil {
			chunkData.Close()
			return fmt.Errorf("couldn't write chunk %d: %v", i, err)
		}
		chunkData.Close()
		os.Remove(chunkPath)
	}
	return nil
}
