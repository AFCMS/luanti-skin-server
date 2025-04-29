package middleware

import (
	"luanti-skin-server/backend/models"

	"github.com/gofiber/fiber/v3"
)

// PermissionHandler Refuse interaction if user doens't have required permission level
//
// Must be used after the auth handler
func PermissionHandler(level int8) fiber.Handler {
	return func(c fiber.Ctx) error {
		user := c.Locals("user").(models.Account)

		if user.PermissionLevel >= level {
			return c.Next()
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message":  "Permission level too low",
				"current":  user.PermissionLevel,
				"expected": level,
			})
		}
	}
}
