package middleware

import (
	"MovieBack/internal/types"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func LoginValidator(c *fiber.Ctx) error {

	userData := types.LoginRequest{}
	validate := validator.New()

	err := c.BodyParser(&userData)
	if err != nil {
		return c.Status(400).JSON("error in parsing")
	}

	errCheck := validate.Struct(userData)

	if errCheck != nil {
		return c.Status(400).JSON(fiber.Map{"validation_error": errCheck.Error()})

	}
	return c.Next()
}
