package main

import (
	"embed"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/template/html"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func main() {

	//go:embed views/*
	var viewsfs embed.FS

	engine := html.NewFileSystem(http.FS(viewsfs), ".html")

	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		ServerHeader:  "goshorly",
		Views:         engine,
	})

	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()

	if err != nil {
		log.Fatal(err)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("home", fiber.Map{})
	})

	type EUrl struct {
		URL string `form:"surl"`
	}

	app.Get("/:id", func(c *fiber.Ctx) error {
		val, err := client.Get(c.Params("id")).Result()
		if err != nil {
			return c.Render("404", fiber.Map{
				"BASEURL": c.Protocol() + "://" + c.Hostname(),
			})
		}
		return c.Redirect(val)
	})

	app.Use(limiter.New(limiter.Config{
		Max:        10,
		Expiration: 60 * time.Second,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Render("home", fiber.Map{
				"ERR": "You have reached the limit of requests! Please check back later. (1 minute)",
			})
		},
	}))

	app.Post("/", func(c *fiber.Ctx) error {
		u := new(EUrl)
		if err := c.BodyParser(u); err != nil {
			return c.Render("home", fiber.Map{
				"ERR": "Parsing Error",
			})
		}

		if !regexp.MustCompile(`^(http|https|mailto|ts3server)://`).MatchString(u.URL) {
			return c.Render("home", fiber.Map{
				"ERR": "Invalid URL, please check and try again.",
			})
		}

		id, err := gonanoid.New(8)

		if err != nil {
			return c.Render("home", fiber.Map{
				"ERR": err.Error(),
			})
		}

		err = client.Set(id, u.URL, 1296000*time.Second).Err()

		if err != nil {
			return c.Render("home", fiber.Map{
				"ERR": err.Error(),
			})
		}

		fURL := c.Protocol() + "://" + c.Hostname() + "/" + id

		return c.Render("home", fiber.Map{
			"URL": fURL,
		})
	})

	log.Fatal(app.Listen(":3000"))
}
