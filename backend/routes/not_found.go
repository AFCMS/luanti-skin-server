package routes

import "github.com/gofiber/fiber/v3"

func NotFound(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotFound)
}
