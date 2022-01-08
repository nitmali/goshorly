package routes

import (
	"git.ucode.space/Phil/goshorly/db"
	"git.ucode.space/Phil/goshorly/utils"
	"github.com/gofiber/fiber/v2"
)

func Gethome(c *fiber.Ctx) error {
	return c.Render("views/home", fiber.Map{
		"GitCommitShort": utils.GitCommitShort,
		"GitBranch":      utils.GitBranch,
		"GitBuild":       utils.GitBuild,
		"TotalLinks":     db.GetTotalLinks(),
		"TotalViews":     db.GetTotalViews(),
	})
}
