package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sayutizxc/klikmanga-scraper/model"
	"github.com/sayutizxc/klikmanga-scraper/service"
	"strconv"
)

func SearchComicController(c *fiber.Ctx) error {
	page := c.Params("page", "1")
	limit := c.Query("limit", "32")
	s := c.Query("s")
	if s == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Status:  fiber.StatusBadRequest,
			Message: "s query parameter is required",
			Data:    nil,
		})
	}
	pageInt, pageErr := strconv.Atoi(page)
	limitInt, limitErr := strconv.Atoi(limit)
	if pageErr != nil || limitErr != nil || pageInt < 1 || limitInt < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Status:  fiber.StatusBadRequest,
			Message: "BAD REQUEST",
			Data:    nil,
		})
	}
	page = strconv.Itoa(pageInt - 1)

	comics, err := service.SearchComicService(s, page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Status:  fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(model.Response{
		Status: fiber.StatusOK,
		Data:   comics,
	})
}
