package routes

import (
	"github.com/gofiber/fiber/v3"

	"luanti-skin-server/backend/database"
	"luanti-skin-server/backend/utils"
)

func Info(c fiber.Ctx) error {
	accountCount, err := database.AccountCount()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	skinCount, err := database.SkinCount()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Return supported OAuth providers
	// Used by the frontend to determine which OAuth buttons to display
	var supportedOAuthProviders []string

	if utils.ConfigOAuthContentDB {
		supportedOAuthProviders = append(supportedOAuthProviders, "contentdb")
	}

	if utils.ConfigOAuthGitHub {
		supportedOAuthProviders = append(supportedOAuthProviders, "github")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"version":                   "1.0",
		"account_count":             accountCount,
		"skin_count":                skinCount,
		"supported_oauth_providers": supportedOAuthProviders,
	})
}
