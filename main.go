package main

import (
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
	engine := html.NewFileSystem(http.Dir("./views"), ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	client := redis.NewClient(&redis.Options{
		Addr:     "db:6379",
		Password: "abc123",
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
		URL string `json:"surl" xml:"surl" form:"surl"`
	}

	app.Get("/:id", func(c *fiber.Ctx) error {
		val, err := client.Get(c.Params("id")).Result()
		if err != nil {
			return c.SendString(err.Error())
		}
		return c.Redirect(val)
	})

	app.Use(limiter.New(limiter.Config{
		Max:        10,
		Expiration: 60 * time.Second,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Render("slowdown", fiber.Map{})
		},
	}))

	app.Post("/", func(c *fiber.Ctx) error {
		u := new(EUrl)
		if err := c.BodyParser(u); err != nil {
			return err
		}

		if !regexp.MustCompile(`^(http|https)://`).MatchString(u.URL) {
			c.Status(http.StatusBadRequest).SendString("Invalid URL")
		}

		id, err := gonanoid.New(8)

		if err != nil {
			c.SendString(err.Error())
		}

		err = client.Set(id, u.URL, 1296000*time.Second).Err()

		if err != nil {
			c.SendString(err.Error())
		}

		return c.SendString("https://" + c.Hostname() + "/" + id)
	})

	log.Fatal(app.Listen(":3000"))
}
