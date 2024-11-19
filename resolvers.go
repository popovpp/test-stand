package main

import (
	"test-stand/schemas"

	"github.com/gofiber/fiber/v2"
)

// GetItem godoc
// @Summary Get an item
// @Description Get an item by its ID
// @ID get-item-by-int
// @Accept  json
// @Produce  json
// @Tags Item
// @Param id path int true "Item ID"
// @Success 200 {object} schemas.Item
// @Failure 400 {object} schemas.HTTPError
// @Failure 404 {object} schemas.HTTPError
// @Failure 500 {object} schemas.HTTPError
// @Router /item/{id} [get]
func GetItem(c *fiber.Ctx) error {
	// Create new Item and returns it
	return c.JSON(schemas.Item{
		Id: c.Params("id"),
	})
}
