package main

import (
	"embed"
	"log"
	"net/http"

	"git.ucode.space/Phil/goshorly/db"
	"git.ucode.space/Phil/goshorly/routes"
	"git.ucode.space/Phil/goshorly/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/template/html"
)

//go:embed views/*
var viewsfs embed.FS

func main() {

	utils.Init_env_vars()
	utils.Init_build_vars()

	db.Init_redis()

	engine := html.NewFileSystem(http.FS(viewsfs), ".html")

	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		Views:         engine,
	})

	app.Get("/", routes.Gethome)

	if utils.ESTATS == "true" {
		app.Get("/stats", routes.GetStats)
	}

	app.Get("/:id", routes.ID)

	app.Use(limiter.New(utils.ConfigLimiter))

	app.Post("/", routes.Posthome)

	log.Fatal(app.Listen(":" + utils.PORT))
}
