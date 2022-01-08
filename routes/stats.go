package routes

import (
	"time"

	"git.ucode.space/Phil/goshorly/db"
	"github.com/gofiber/fiber/v2"
)

func GetStats(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"Timestamp":  time.Now(),
		"Totallinks": db.GetTotalLinks(),
		"Totalviews": db.GetTotalViews(),
	})
}
