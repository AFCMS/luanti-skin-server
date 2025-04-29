package routes

import (
	"luanti-skin-server/backend/database"
	"luanti-skin-server/backend/models"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

// SkinDetails Return the skin file
func SkinDetails(c fiber.Ctx) error {
	var skin models.Skin

	id, err := uuid.Parse(c.Params("uuid"))
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	skin, err = database.SkinFromUUID(id)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(skin)
}
