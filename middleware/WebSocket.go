package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/MertJSX/folder-host-go/database/logs"
	"github.com/MertJSX/folder-host-go/database/users"
	"github.com/MertJSX/folder-host-go/types"
	"github.com/MertJSX/folder-host-go/utils"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func WsConnect(c *fiber.Ctx) error {
	if !websocket.IsWebSocketUpgrade(c) {
		return c.Next()
	}

	token := c.Query("token")
	if token == "" {
		authHeader := c.Get("Authorization")
		if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
			token = strings.TrimPrefix(authHeader, "Bearer ")
		}
	}

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token required for WebSocket connection",
		})
	}

	username, err := utils.VerifyToken(token, utils.Config.SecretJwtKey)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "invalid token"})
	}

	foundAccount, err := users.GetUserByUsername(username)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"err": "account not found"})
	}

	c.Locals("username", foundAccount.Username)
	c.Locals("account", foundAccount)
	c.Locals("token", token)
	c.Locals("allowed", true)

	return c.Next()
}

func HandleWebsocket(c *websocket.Conn) {
	var path string = c.Params("path")

	path, err := url.PathUnescape(path)

	if err != nil {
		c.Close()
		return
	}

	config := &utils.Config
	path = config.Folder + path

	utils.AddClient(c, path)
	updateClientsCount(path)

	defer updateClientsCount(path)
	defer utils.RemoveClient(c)

	var username string = c.Locals("username").(string)

	defer utils.TriggerPendingLog(username, path)

	// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
	var (
		mt  int
		msg []byte
	)

	for {
		if mt, msg, err = c.ReadMessage(); err != nil {
			log.Println("WebSocket read error:", err)
			break
		}

		if err := processWebSocketMessage(msg, path, c, mt); err != nil {
			log.Println("Message processing error:", err)
		}
	}

	log.Printf("WebSocket disconnected - User: %s, Path: %s", username, path)
	c.Close()
}

func updateClientsCount(path string) {
	clientsCount, err := json.Marshal(fiber.Map{
		"type":  "editor-update-usercount",
		"count": utils.GetClientsCount(path),
	})

	if err == nil {
		utils.SendToAll(path, 1, clientsCount)
	}
}

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

		handleUnzip(c, mt, message)
	default:
		log.Printf("Unknown message type: %s\n", message.Type)
	}

	return nil
}

func handleUnzip(c *websocket.Conn, mt int, message types.EditorChange) {
	src := utils.Config.Folder + message.Path
	dest := fmt.Sprintf("%s%s/%s", utils.Config.Folder, utils.GetParentPath(message.Path), utils.GetPureFileName(message.Path))

	for index := 1; utils.IsExistingPath(dest); index++ {
		dest = fmt.Sprintf("%s (%d)", dest, index)
	}

	utils.Unzip(src, dest, func(totalSize int64, isCompleted bool, abortMsg string) {
		unzipProgress, _ := json.Marshal(fiber.Map{
			"type":        "unzip-progress",
			"totalSize":   utils.ConvertBytesToString(totalSize),
			"isCompleted": isCompleted,
			"abortMsg":    abortMsg,
		})

		c.WriteMessage(mt, unzipProgress)
	})
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
	return os.WriteFile(filepath, []byte(content), 0644)
}
