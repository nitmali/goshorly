package routes

import (
	"regexp"
	"time"

	"git.ucode.space/Phil/goshorly/db"
	"git.ucode.space/Phil/goshorly/utils"
	"github.com/gofiber/fiber/v2"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type eurl struct {
	URL string `form:"surl"`
}

func Posthome(c *fiber.Ctx) error {
	u := new(eurl)
	if err := c.BodyParser(u); err != nil {
		return c.Render("views/home", fiber.Map{
			"ERR":            "Parsing Error",
			"GitCommitShort": utils.GitCommitShort,
			"GitBranch":      utils.GitBranch,
			"GitBuild":       utils.GitBuild,
		})
	}

	if !regexp.MustCompile(`^(http|https|mailto|ts3server)://`).MatchString(u.URL) {
		return c.Render("views/home", fiber.Map{
			"ERR":            "Invalid URL, please check and try again.",
			"GitCommitShort": utils.GitCommitShort,
			"GitBranch":      utils.GitBranch,
			"GitBuild":       utils.GitBuild,
		})
	}

	id, err := gonanoid.New(8)

	if err != nil {
		return c.Render("views/home", fiber.Map{
			"ERR":            err.Error(),
			"GitCommitShort": utils.GitCommitShort,
			"GitBranch":      utils.GitBranch,
			"GitBuild":       utils.GitBuild,
		})
	}

	err = db.Client.Set(id, u.URL, 1296000*time.Second).Err()

	if err != nil {
		return c.Render("views/home", fiber.Map{
			"ERR":            err.Error(),
			"GitCommitShort": utils.GitCommitShort,
			"GitBranch":      utils.GitBranch,
			"GitBuild":       utils.GitBuild,
		})
	}

	fURL := utils.URL + id

	return c.Render("views/home", fiber.Map{
		"URL":            fURL,
		"GitCommitShort": utils.GitCommitShort,
		"GitBranch":      utils.GitBranch,
		"GitBuild":       utils.GitBuild,
	})
}
