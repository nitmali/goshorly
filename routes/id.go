package routes

import (
	"git.ucode.space/Phil/goshorly/db"
	"git.ucode.space/Phil/goshorly/utils"
	"github.com/gofiber/fiber/v2"
)

func ID(c *fiber.Ctx) error {
	val, err := db.Client.Get(c.Params("id")).Result()
	if err != nil {
		return c.Render("views/404", fiber.Map{
			"BASEURL": utils.URL,
		})
	}
	return c.Redirect(val)
}
