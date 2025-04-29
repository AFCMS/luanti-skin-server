package routes

import (
	"luanti-skin-server/backend/database"

	"github.com/gofiber/fiber/v3"
)

func UsersList(c fiber.Ctx) error {
	result, err := database.AccountList()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
