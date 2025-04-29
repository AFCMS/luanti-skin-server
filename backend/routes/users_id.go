package routes

import (
	"luanti-skin-server/backend/database"
	"luanti-skin-server/backend/models"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

// UsersID Return the skin file
func UsersID(c fiber.Ctx) error {
	var a models.Account

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	a, err = database.AccountFromID(uint(id))
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(a)
}
