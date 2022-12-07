package controllers

import (
	"github.com/LulianoM/api-geoprocess/src/database"
	"github.com/LulianoM/api-geoprocess/src/models"
	"github.com/gofiber/fiber/v2"
)

func CommercialsCreateNew(c *fiber.Ctx) error {
	var commercial models.Commercials
	if err := c.BodyParser(&commercial); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"error": err,
		})
	}

	if result := database.DB.Create(&commercial); result.Error != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"error": result.Error,
		})
	}

	return c.JSON(&commercial)
}

func CommercialsGetAll(c *fiber.Ctx) error {
	var commercials []models.Commercials

	if result := database.DB.Find(&commercials); result.Error != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"error": result.Error,
		})
	}

	return c.JSON(fiber.Map{
		"data": commercials,
	})
}
