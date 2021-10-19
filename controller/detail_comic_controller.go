package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sayutizxc/klikmanga-scraper/model"
	"github.com/sayutizxc/klikmanga-scraper/service"
)

func DetailComicController(c *fiber.Ctx) error {
	url := c.Query("url", "")
	detailComic, err := service.DetailComicService(url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Status:  fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(model.Response{
		Status: fiber.StatusOK,
		Data:   detailComic,
	})
}
