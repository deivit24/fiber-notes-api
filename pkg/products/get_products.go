package products

import (
	"github.com/deivit24/api-db/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

func (h handler) GetProducts(c *fiber.Ctx) error {
	var products []models.Product

	if result := h.DB.Find(&products); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&products)
}
