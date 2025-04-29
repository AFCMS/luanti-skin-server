package routes

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"

	"luanti-skin-server/backend/database"
	"luanti-skin-server/backend/models"
	"luanti-skin-server/backend/types"
)

func AccountRegister(c fiber.Ctx) error {
	input := new(types.InputRegister)

	if err := c.Bind().JSON(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error on register request", "data": err.Error()})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	user := models.Account{
		Username:  input.Username,
		Password:  password,
		CreatedAt: time.Now(),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error on register request", "data": err.Error()})
	}

	return c.JSON(types.OutputRegister{
		Id:           user.ID,
		Username:     user.Username,
		CreationDate: user.CreatedAt.Unix(),
	})
}
