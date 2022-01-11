package routes

import (
	"log"
	"regexp"
	"time"

	"git.ucode.space/Phil/goshorly/db"
	"git.ucode.space/Phil/goshorly/utils"
	"github.com/gofiber/fiber/v2"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type eurl struct {
	CLI bool   `json:"cli"`
	URL string `json:"surl" form:"surl"`
}

func Posthome(c *fiber.Ctx) error {
	u := new(eurl)
	if err := c.BodyParser(u); err != nil {

		if u.CLI {
			return c.Status(500).JSON(&fiber.Map{
				"success": false,
				"msg":     "Parsing Error",
			})
		}

		return c.Status(500).Render("views/home", fiber.Map{
			"ERR":                 "Parsing Error",
			"CI_COMMIT_SHORT_SHA": utils.CI_COMMIT_SHORT_SHA,
			"CI_COMMIT_BRANCH":    utils.CI_COMMIT_BRANCH,
			"CI_COMMIT_TAG":       utils.CI_COMMIT_TAG,
			"CI_TAGGED":           utils.CI_TAGGED,
			"CI_BUILD":            utils.CI_BUILD,
			"TotalLinks":          db.GetTotalLinks(),
			"TotalViews":          db.GetTotalViews(),
		})
	}

	if !regexp.MustCompile(`^(http|https|mailto|ts3server)://`).MatchString(u.URL) {

		if u.CLI {
			return c.Status(424).JSON(&fiber.Map{
				"success": false,
				"msg":     "Invalid URL",
			})
		}

		return c.Status(424).Render("views/home", fiber.Map{
			"ERR":                 "Invalid URL, please check and try again.",
			"CI_COMMIT_SHORT_SHA": utils.CI_COMMIT_SHORT_SHA,
			"CI_COMMIT_BRANCH":    utils.CI_COMMIT_BRANCH,
			"CI_COMMIT_TAG":       utils.CI_COMMIT_TAG,
			"CI_TAGGED":           utils.CI_TAGGED,
			"CI_BUILD":            utils.CI_BUILD,
			"TotalLinks":          db.GetTotalLinks(),
			"TotalViews":          db.GetTotalViews(),
		})
	}

	id, err := gonanoid.New(8)

	if err != nil {

		if u.CLI {
			return c.Status(500).JSON(&fiber.Map{
				"success": false,
				"msg":     err.Error(),
			})
		}

		return c.Status(500).Render("views/home", fiber.Map{
			"ERR":                 err.Error(),
			"CI_COMMIT_SHORT_SHA": utils.CI_COMMIT_SHORT_SHA,
			"CI_COMMIT_BRANCH":    utils.CI_COMMIT_BRANCH,
			"CI_COMMIT_TAG":       utils.CI_COMMIT_TAG,
			"CI_TAGGED":           utils.CI_TAGGED,
			"CI_BUILD":            utils.CI_BUILD,
			"TotalLinks":          db.GetTotalLinks(),
			"TotalViews":          db.GetTotalViews(),
		})
	}

	err = db.Client.Set(id, u.URL, 1296000*time.Second).Err()

	if err != nil {
		if u.CLI {
			return c.Status(500).JSON(&fiber.Map{
				"success": false,
				"msg":     err.Error(),
			})
		}

		return c.Status(500).Render("views/home", fiber.Map{
			"ERR":                 err.Error(),
			"CI_COMMIT_SHORT_SHA": utils.CI_COMMIT_SHORT_SHA,
			"CI_COMMIT_BRANCH":    utils.CI_COMMIT_BRANCH,
			"CI_COMMIT_TAG":       utils.CI_COMMIT_TAG,
			"CI_TAGGED":           utils.CI_TAGGED,
			"CI_BUILD":            utils.CI_BUILD,
			"TotalLinks":          db.GetTotalLinks(),
			"TotalViews":          db.GetTotalViews(),
		})
	}

	fURL := utils.URL + id

	// Increase Total Links into the DB
	_, err = db.StatsIncreaseTotalLinks()

	if err != nil {
		log.Fatalln(err.Error())
	}

	if u.CLI {
		return c.Status(201).JSON(&fiber.Map{
			"success": true,
			"URL":     fURL,
		})
	}

	return c.Status(201).Render("views/home", fiber.Map{
		"URL":                 fURL,
		"CI_COMMIT_SHORT_SHA": utils.CI_COMMIT_SHORT_SHA,
		"CI_COMMIT_BRANCH":    utils.CI_COMMIT_BRANCH,
		"CI_COMMIT_TAG":       utils.CI_COMMIT_TAG,
		"CI_TAGGED":           utils.CI_TAGGED,
		"CI_BUILD":            utils.CI_BUILD,
		"TotalLinks":          db.GetTotalLinks(),
		"TotalViews":          db.GetTotalViews(),
	})
}
