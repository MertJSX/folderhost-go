package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"unicode/utf8"

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
		return c.Status(401).JSON(fiber.Map{"error": "Invalid token"})
	}

	config := &utils.Config
	var accountFound bool
	for _, v := range config.Accounts {
		if v.Name == username {
			c.Locals("username", username)
			c.Locals("account", v)
			c.Locals("token", token)
			c.Locals("allowed", true)
			accountFound = true
			break
		}
	}

	if !accountFound {
		return c.Status(401).JSON(fiber.Map{"error": "Account not found"})
	}

	var path string = c.Query("path")

	fmt.Printf("Path: %s\n", path)

	decodedPath, err := url.PathUnescape(path)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid path"})
	}

	fmt.Printf("Decoded path: %s\n", decodedPath)

	c.Locals("path", decodedPath)

	return c.Next()
}

func HandleWebsocket(c *websocket.Conn) {
	//log.Println(c.Locals("allowed")) // true
	var path string = c.Params("path")

	fmt.Printf("Path: %s\n", path)

	path, _ = url.PathUnescape(path)
	config := &utils.Config
	path = fmt.Sprintf("%s%s", config.Folder, path)

	utils.AddClient(c, path)

	defer utils.RemoveClient(c)

	var username string = c.Locals("username").(string)

	// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
	var (
		mt  int
		msg []byte
		err error
	)

	for {
		if mt, msg, err = c.ReadMessage(); err != nil {
			log.Println("WebSocket read error:", err)
			break
		}

		utils.SendToAll(path, mt, msg, c)

		if err := processWebSocketMessage(msg, path, c); err != nil {
			log.Println("Message processing error:", err)
		}
	}

	log.Printf("WebSocket disconnected - User: %s, Path: %s", username, path)
	c.Close()
}

func processWebSocketMessage(msg []byte, filePath string, c *websocket.Conn) error {
	var message types.EditorChange

	if err := json.Unmarshal(msg, &message); err != nil {
		return err
	}

	fmt.Printf("Change: %s\n", message.Change.Text)

	fmt.Printf("Filepath: %s\n", filePath)

	switch message.Type {
	case "editor-change":
		return applyEditorChange(filePath, message.Change)
	default:
		log.Printf("Unknown message type: %s", message.Type)
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
	return os.WriteFile(filepath, []byte(content), 0644)
}
