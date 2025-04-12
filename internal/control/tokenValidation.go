package control

import (
	"fmt"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"

	"github.com/gofiber/fiber/v2"
)

func TokenValidation(c *fiber.Ctx) error {
	// Get the Authorization header
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return c.Status(401).JSON(fiber.Map{
			"error": "Authorization header missing",
		})
	}

	token := strings.Split(authHeader, " ")[1]

	validUser, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(os.Getenv("JWT_PASSWORD")), nil
	})
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "verification error",
		})
	}
	fmt.Print(validUser)
	return c.Status(200).JSON(fiber.Map{"token": token})
	//  else {
	// 	fmt.Println("authHeader", authHeader)
	// 	return c.Status(200).JSON("success")
	// }

}
