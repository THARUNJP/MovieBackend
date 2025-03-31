package control

import (
	"MovieBack/config"
	"MovieBack/internal/custom"

	"github.com/gofiber/fiber/v2"
)

func GetMovies(c *fiber.Ctx) error {

	getData, err := config.GetRecords("SELECT * FROM movies", nil)
	if err != nil {
		return c.SendStatus(fiber.ErrBadGateway.Code)
	} else {
		custom.Map(getData)
		return c.Status(200).JSON(getData)
	}

}
