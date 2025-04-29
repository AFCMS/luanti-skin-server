package routes

import (
	"luanti-skin-server/backend/database"
	"luanti-skin-server/backend/models"
	"luanti-skin-server/backend/types"

	"github.com/gofiber/fiber/v3"
)

func SkinList(c fiber.Ctx) error {
	// Parse query
	queryR := new(types.QuerySkinList)
	if err := c.Bind().Query(queryR); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Bad request")
	}

	// Convert values to handle GORM API
	count := int(queryR.Count)

	if queryR.Count == 0 {
		count = 10
	}

	page := int(queryR.Page)

	if queryR.Page == 0 {
		page = 0
	}

	// Query database
	var result []models.Skin

	if err := database.DB.Limit(count).Offset(page * count).Find(&result).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Cannot interact with database")
	}

	return c.JSON(result)
}
