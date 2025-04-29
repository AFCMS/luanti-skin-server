package routes

import (
	"luanti-skin-server/backend/models"

	"github.com/gofiber/fiber/v3"
)

func AccountUser(c fiber.Ctx) error {
	user := c.Locals("user").(models.Account)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":               user.ID,
		"username":         user.Username,
		"permission_level": user.PermissionLevel,
		"cdb_username":     user.CDBUsername,
	})
}
