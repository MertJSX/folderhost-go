package websocket

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/MertJSX/folder-host-go/database/logs"
	"github.com/MertJSX/folder-host-go/types"
	"github.com/MertJSX/folder-host-go/utils"
	"github.com/MertJSX/folder-host-go/utils/cache"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func processWebSocketMessage(msg []byte, filePath string, c *websocket.Conn, mt int) error {
	var message types.EditorChange
	var account types.Account = c.Locals("account").(types.Account)

	if err := json.Unmarshal(msg, &message); err != nil {
		return err
	}

	switch message.Type {
	case "editor-change":
		if !account.Permissions.Change {
			permissionError, _ := json.Marshal(fiber.Map{
				"type":  "error",
				"error": "You don't have permission to change!",
			})

			c.WriteMessage(mt, permissionError)
			return nil // Server doesn't care about permission errors
		}

		utils.ScheduleDebouncedLog(account.Username, filePath)

		utils.SendToAllExclude(filePath, mt, msg, c)
		return applyEditorChange(filePath, message.Change)
	case "unzip":
		if !account.Permissions.Extract {
			permissionError, _ := json.Marshal(fiber.Map{
				"type":  "error",
				"error": "You don't have permission to unzip!",
			})

			c.WriteMessage(mt, permissionError)
			return nil // Server doesn't care about permission errors
		}

		logs.CreateLog(types.AuditLog{
			Username:    account.Username,
			Action:      "Extract file",
			Description: fmt.Sprintf("%s started unzipping %s file.", account.Username, message.Path),
		})

		HandleUnzip(c, mt, message)
	}

	return nil
}

func applyEditorChange(filePath string, change types.ChangeData) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	lines := strings.Split(string(content), "\n")

	startLine := change.Range.StartLineNumber - 1
	startCol := change.Range.StartColumn - 1
	endLine := change.Range.EndLineNumber - 1
	endCol := change.Range.EndColumn - 1

	startCol = adjustUTF8Position(lines, startLine, startCol)
	endCol = adjustUTF8Position(lines, endLine, endCol)

	switch change.Type {
	case "insert":
		return applyInsert(filePath, lines, startLine, startCol, change.Text)
	case "delete":
		return applyDelete(filePath, lines, startLine, startCol, endLine, endCol)
	case "replace":
		return applyReplace(filePath, lines, startLine, startCol, endLine, endCol, change.Text)
	default:
		return nil
	}
}

func adjustUTF8Position(lines []string, lineNum, col int) int {
	if lineNum < 0 || lineNum >= len(lines) {
		return col
	}

	line := lines[lineNum]

	if col > utf8.RuneCountInString(line) {
		return utf8.RuneCountInString(line)
	}

	return col
}

func applyInsert(filePath string, lines []string, lineNum, col int, text string) error {
	if lineNum < 0 || lineNum >= len(lines) {
		return fmt.Errorf("line number out of range: %d", lineNum)
	}

	line := lines[lineNum]
	runes := []rune(line)

	if col < 0 {
		col = 0
	}
	if col > len(runes) {
		col = len(runes)
	}

	newRunes := make([]rune, 0, len(runes)+len([]rune(text)))
	newRunes = append(newRunes, runes[:col]...)
	newRunes = append(newRunes, []rune(text)...)
	newRunes = append(newRunes, runes[col:]...)

	lines[lineNum] = string(newRunes)
	return writeFile(filePath, lines)
}

func applyDelete(filePath string, lines []string, startLine, startCol, endLine, endCol int) error {
	if startLine < 0 || startLine >= len(lines) || endLine < 0 || endLine >= len(lines) {
		return fmt.Errorf("line numbers out of range: %d-%d", startLine, endLine)
	}

	if startLine == endLine {
		line := lines[startLine]
		runes := []rune(line)

		startCol = utils.Clamp(startCol, 0, len(runes))
		endCol = utils.Clamp(endCol, 0, len(runes))

		if startCol >= endCol {
			return nil
		}

		newRunes := make([]rune, 0, len(runes)-(endCol-startCol))
		newRunes = append(newRunes, runes[:startCol]...)
		newRunes = append(newRunes, runes[endCol:]...)
		lines[startLine] = string(newRunes)

	} else {
		firstLineRunes := []rune(lines[startLine])
		lastLineRunes := []rune(lines[endLine])

		startCol = utils.Clamp(startCol, 0, len(firstLineRunes))
		endCol = utils.Clamp(endCol, 0, len(lastLineRunes))

		newFirstLine := string(firstLineRunes[:startCol])

		newLastLine := string(lastLineRunes[endCol:])

		lines[startLine] = newFirstLine + newLastLine

		if startLine < endLine {
			lines = append(lines[:startLine+1], lines[endLine+1:]...)
		}
	}

	return writeFile(filePath, lines)
}

func applyReplace(filePath string, lines []string, startLine, startCol, endLine, endCol int, text string) error {
	if startLine == endLine {
		line := lines[startLine]
		runes := []rune(line)

		startCol = utils.Clamp(startCol, 0, len(runes))
		endCol = utils.Clamp(endCol, 0, len(runes))

		if startCol > endCol {
			return fmt.Errorf("invalid range: startCol > endCol")
		}

		textRunes := []rune(text)
		newRunes := make([]rune, 0, len(runes)-(endCol-startCol)+len(textRunes))
		newRunes = append(newRunes, runes[:startCol]...)
		newRunes = append(newRunes, textRunes...)
		newRunes = append(newRunes, runes[endCol:]...)
		lines[startLine] = string(newRunes)

	} else {
		firstLineRunes := []rune(lines[startLine])
		startCol = utils.Clamp(startCol, 0, len(firstLineRunes))

		firstPart := string(firstLineRunes[:startCol])
		newFirstLine := firstPart + text

		lastLineRunes := []rune(lines[endLine])
		endCol = utils.Clamp(endCol, 0, len(lastLineRunes))
		lastPart := string(lastLineRunes[endCol:])

		lines[startLine] = newFirstLine + lastPart

		if startLine < endLine {
			lines = append(lines[:startLine+1], lines[endLine+1:]...)
		}
	}

	return writeFile(filePath, lines)
}

func writeFile(filepath string, lines []string) error {
	content := strings.Join(lines, "\n")

	watcherCache, ok := cache.EditorWatcherCache.Get(filepath)
	if !ok {
		return fmt.Errorf("cannot get watcher cache")
	}

	watcherCache.IsWriting = true

	cache.EditorWatcherCache.SetWithoutTTL(filepath, watcherCache)

	err := os.WriteFile(filepath, []byte(content), 0644)

	if err != nil {
		return err
	}

	fileStat, err := os.Stat(filepath)

	if err != nil {
		return err
	}

	watcherCache.LastModTime = fileStat.ModTime()
	watcherCache.IsWriting = false

	cache.EditorWatcherCache.SetWithoutTTL(filepath, watcherCache)

	if directoryCache, ok := cache.DirectoryCache.Get(utils.GetParentPath(filepath) + "/"); ok {
		for index, v := range directoryCache.Items {
			v.SizeBytes = fileStat.Size()
			directoryCache.Items[index] = v
		}
		cache.DirectoryCache.Set(utils.GetParentPath(filepath)+"/", directoryCache, 600)
	}

	return nil
}
