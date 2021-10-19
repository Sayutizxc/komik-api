package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sayutizxc/klikmanga-scraper/controller"
	"log"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	comic := app.Group("/api/comic")
	comic.Get("/home/:page?", controller.ListComicController)
	comic.Get("/detail", controller.DetailComicController)
	comic.Get("/chapter", controller.ChapterImagesController)
	comic.Get("/search/:page?", controller.SearchComicController)

	// Start server
	log.Fatal(app.Listen(":3000"))
}
