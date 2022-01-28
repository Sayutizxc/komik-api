package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sayutizxc/klikmanga-scraper/model"
	"github.com/sayutizxc/klikmanga-scraper/scraper"
)

func ChapterImagesController(c *fiber.Ctx) error {
	url := c.Query("url", "")
	chaperImages, err := scraper.ChapterImagesScraper(url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Status:  fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if chaperImages.Images == nil && chaperImages.Chapter == "" {
		return c.Status(fiber.StatusNotFound).JSON(model.Response{
			Status:  fiber.StatusNotFound,
			Message: "Not Found",
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.Response{
		Status:  fiber.StatusOK,
		Message: "Ok",
		Data:    chaperImages,
	})
}
