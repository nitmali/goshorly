package routes

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"time"

	"git.ucode.space/Phil/goshorly/db"
	"git.ucode.space/Phil/goshorly/utils"
)

func ID(c *fiber.Ctx) error {
	val, err := db.Get(c.Params("id"))

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
	// reset effective time
	err = db.SetEX(c.Params("id"), val, 7*24*time.Minute)
	if err != nil {
		log.Fatal(err.Error())
	}
	return c.Redirect(val)
}
