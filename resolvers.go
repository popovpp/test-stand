package main

import (
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
// @Success 200 {object} Item
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /item/{id} [get]
func GetItem(c *fiber.Ctx) error {
	// Create new Item and returns it
	return c.JSON(Item{
		Id: c.Params("id"),
	})
}

type Item struct {
	Id string
}

type HTTPError struct {
	Status  string
	Message string
}
