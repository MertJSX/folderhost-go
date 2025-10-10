package routes

import (
	"strconv"

	"github.com/MertJSX/folder-host-go/database/logs"
	"github.com/MertJSX/folder-host-go/types"
	"github.com/gofiber/fiber/v2"
)

func Logs(c *fiber.Ctx) error {
	if !c.Locals("account").(types.Account).Permissions.ReadLogs {
		return c.Status(403).JSON(
			fiber.Map{"err": "No permission!"},
		)
	}

	page := c.Query("page")

	if page == "" {
		return c.Status(400).JSON(fiber.Map{
			"err": "Page parameter is required",
		})
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"err": "Page parameter must be a valid integer",
		})
	}

	pageInt--

	if pageInt < 0 {
		return c.Status(400).JSON(fiber.Map{
			"err": "Page parameter cannot be negative",
		})
	}

	// My logic: If pageInt is 0, we will skip 0 foundLogs. That means that we wil get the last 10 foundLogs.
	// If pageInt is 1 we will skip 10 foundLogs.
	foundLogs, err := logs.SearchLogs(20, 20*pageInt)
	if err != nil {
		return c.Status(500).JSON(
			fiber.Map{"err": "Unknown error!"},
		)
	}
	pageInt++
	nextLogs, err := logs.SearchLogs(20, 20*pageInt)
	if err != nil {
		return c.Status(500).JSON(
			fiber.Map{"err": "Unknown error!"},
		)
	}

	return c.Status(200).JSON(
		fiber.Map{
			"logs":   foundLogs,
			"isLast": len(nextLogs) == 0,
		},
	)
}
