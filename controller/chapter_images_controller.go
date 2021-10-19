package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sayutizxc/klikmanga-scraper/model"
	"github.com/sayutizxc/klikmanga-scraper/service"
)

func ChapterImagesController(c *fiber.Ctx) error {
	url := c.Query("url", "")
	chaperImages, err := service.ChapterImagesService(url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Status:  fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(model.Response{
		Status: fiber.StatusOK,
		Data:   chaperImages,
	})
}
