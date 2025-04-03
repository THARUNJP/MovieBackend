package middleware

import "github.com/gofiber/fiber/v2"

func TokenValidation(c *fiber.Ctx) error {
	return c.Next()

}
