package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sayutizxc/klikmanga-scraper/model"
	"github.com/sayutizxc/klikmanga-scraper/scraper"
)

func DetailComicController(c *fiber.Ctx) error {
	url := c.Query("url", "")
	detailComic, err := scraper.DetailComicScraper(url)
	status := fiber.StatusOK
	if err != nil {
		if err.Error() == "Not Found" {
			status = fiber.StatusNotFound
		} else {
			status = fiber.StatusInternalServerError
		}
		return c.Status(status).JSON(model.Response{
			Status:  status,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.Response{
		Status:  fiber.StatusOK,
		Message: "Ok",
		Data:    detailComic,
	})
}
