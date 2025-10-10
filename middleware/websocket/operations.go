package websocket

import (
	"encoding/json"
	"fmt"

	"github.com/MertJSX/folder-host-go/types"
	"github.com/MertJSX/folder-host-go/utils"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func HandleUnzip(c *websocket.Conn, mt int, message types.EditorChange) {
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
