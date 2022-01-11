package routes

import (
	"git.ucode.space/Phil/goshorly/db"
	"git.ucode.space/Phil/goshorly/utils"
	"github.com/gofiber/fiber/v2"
)

func Gethome(c *fiber.Ctx) error {
	return c.Render("views/home", fiber.Map{
		"CI_COMMIT_SHORT_SHA": utils.CI_COMMIT_SHORT_SHA,
		"CI_COMMIT_BRANCH":    utils.CI_COMMIT_BRANCH,
		"CI_COMMIT_TAG":       utils.CI_COMMIT_TAG,
		"CI_TAGGED":           utils.CI_TAGGED,
		"CI_BUILD":            utils.CI_BUILD,
		"TotalLinks":          db.GetTotalLinks(),
		"TotalViews":          db.GetTotalViews(),
	})
}
