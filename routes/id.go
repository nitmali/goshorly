package routes

import (
	"log"

	"git.ucode.space/Phil/goshorly/db"
	"git.ucode.space/Phil/goshorly/utils"
	"github.com/gofiber/fiber/v2"
)

func ID(c *fiber.Ctx) error {
	val, err := db.Client.Get(c.Params("id")).Result()

	if c.Get("CLI") == "1" {
		if err != nil {
			return c.Status(404).JSON(&fiber.Map{
				"error": true,
				"url":   "URL not found",
			})
		} else {
			_, err = db.StatsIncreaseViewsLinks()
			if err != nil {
				log.Fatal(err.Error())
			}
			return c.Status(301).JSON(&fiber.Map{
				"error": false,
				"url":   val,
			})
		}
	}

	if err != nil {
		return c.Render("views/404", fiber.Map{
			"BASEURL": utils.URL,
		})
	}

	_, err = db.StatsIncreaseViewsLinks()
	if err != nil {
		log.Fatal(err.Error())
	}
	return c.Redirect(val)
}
