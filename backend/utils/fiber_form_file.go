package utils

import (
	"mime/multipart"

	"github.com/gofiber/fiber/v3"
)

func LoadFormFile(c fiber.Ctx, name string) ([]byte, error) {
	// Get file header
	var fileHeader *multipart.FileHeader
	var err error
	if fileHeader, err = c.FormFile("data"); err != nil {
		return nil, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error on create request", "data": err})
	}

	// Validate file size
	if fileHeader.Size > 100000 {
		return nil, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error on create request", "data": "Too big file"})
	}

	// Get file
	var file multipart.File
	if file, err = fileHeader.Open(); err != nil {
		return nil, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error on create request", "data": err})
	}

	var b = make([]byte, fileHeader.Size)

	// Read file content
	var _ int
	if _, err = file.Read(b); err != nil {
		return nil, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error on create request", "data": err})
	}

	return b, nil
}
