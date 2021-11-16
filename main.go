package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sayutizxc/klikmanga-scraper/controller"
	"log"
	"os"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Dokumentasi sementara bisa dibaca di : https://github.com/Sayutizxc/klikmanga-api")
	})
	comic := app.Group("/api/comic")
	comic.Get("/home/:page?", controller.ListComicController)
	comic.Get("/detail", controller.DetailComicController)
	comic.Get("/chapter", controller.ChapterImagesController)
	comic.Get("/search/:page?", controller.SearchComicController)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Default port if not specified
	}

	log.Fatal(app.Listen(":" + port))
}
